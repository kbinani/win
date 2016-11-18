package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	. "github.com/kbinani/win/internal"
	"github.com/ogier/pflag"
)

func main() {
	flagMinGWVersion := pflag.String("mingw-version", DefaultMinGWVersion, "Version number of MinGW") // https://github.com/mirror/mingw-w64/releases
	flagWineVersion := pflag.String("wine-version", DefaultWineVersion, "Version number of Wine")     // https://github.com/wine-mirror/wine/releases
	// flagFunctions := pflag.String("functions", "", "Comma separated list of Win32API (ex. --functions=CreateFile,FindWindowEx)")
	// flagAll := pflag.Bool("all", false, "Set this true if you want all Win32API")
	pflag.Parse()

	funcs := getFuncList(*flagMinGWVersion, *flagWineVersion)

	// Print functions for each DLL
	referencedTypes := NewStringSet()
	allExportedFuncs := []Func{}
	for dll, funcList := range funcs {
		types, exportedFuncs, err := printFuncs(funcList, removeFileExt(dll))
		if err != nil {
			panic(err)
		}
		referencedTypes.Merge(types)
		for _, f := range exportedFuncs {
			allExportedFuncs = append(allExportedFuncs, f)
		}
	}

	// Copy referenced types from internal/types/{typename}.go
	if err := createTypedefFile(referencedTypes); err != nil {
		panic(err)
	}

	// Copy extra files
	extraFiles := []string{
		"const.go",
		"win.go",
		"win_windows.go",
		"win_nonwindows.go",
	}
	for _, fileName := range extraFiles {
		src, err := os.Open(filepath.Join("internal", "win", fileName))
		if err != nil {
			panic(err)
		}
		dest, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		io.Copy(dest, src)
		dest.Close()
		src.Close()
	}

	// Report API coverage
	coverage := make(map[string]*Rat)
	for _, funcList := range funcs {
		for _, fun := range funcList {
			exported := false
			for _, f := range allExportedFuncs {
				if fun.Name == f.Name {
					exported = true
					break
				}
			}
			r, ok := coverage[fun.Dll]
			if ok {
				num := r.Num
				if exported {
					num += 1
				}

				denom := r.Denom + 1
				coverage[fun.Dll] = NewRat(num, denom)
			} else {
				if exported {
					coverage[fun.Dll] = NewRat(1, 1)
				} else {
					coverage[fun.Dll] = NewRat(0, 1)
				}
			}
		}
	}
	fmt.Printf("API coverage:\n")
	var coverages APICoverages
	for dll, cov := range coverage {
		coverages = append(coverages, APICoverage{dll, cov})
	}
	sort.Sort(coverages)
	for _, cov := range coverages {
		rate := cov.Coverage.Float64() * 100.0
		fmt.Printf("  %-16s: %4d/%4d (%6.2f)\n", cov.Dll, cov.Coverage.Num, cov.Coverage.Denom, rate)
	}
}

func loadCache(cacheFilePath string) (map[string][]Func, error) {
	funcs := make(map[string][]Func)
	if !IsFileExist(cacheFilePath) {
		return funcs, fmt.Errorf("Cache file not found: %s", cacheFilePath)
	}
	fp, err := os.Open(cacheFilePath)
	if err != nil {
		return funcs, fmt.Errorf("Cannot open cache: %s", cacheFilePath)
	}
	defer fp.Close()
	bytes, e := ioutil.ReadAll(fp)
	if e != nil {
		return funcs, e
	}
	e2 := json.Unmarshal(bytes, &funcs)
	if e2 != nil {
		return make(map[string][]Func), e2
	}
	return funcs, nil
}

func saveCache(cache map[string][]Func, cacheFilePath string) {
	fp, err := os.Create(cacheFilePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	enc := json.NewEncoder(fp)
	enc.Encode(&cache)
}

func getFuncList(minGWVersion string, wineVersion string) map[string][]Func {
	cacheFilePath := filepath.Join("internal", "cache.json")
	cache, err := loadCache(cacheFilePath)
	if err == nil {
		return cache
	}

	// Download mingw-w64-vX.Y.Z.zip from GitHub.
	minGWZipPath, err := downloadMinGWZip(minGWVersion)
	if err != nil {
		panic(err)
	}

	zipFile, err := os.Open(minGWZipPath)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	reader, err := zip.NewReader(zipFile, FileSize(minGWZipPath))
	if err != nil {
		panic(err)
	}

	// Extract Win32 API's from header
	funcsInHeader := []Func{}
	for _, f := range win32headers(reader) {
		r, err := f.Open()
		if err != nil {
			continue
		}
		for _, fun := range extractWinAPIDefinition(r) {
			funcsInHeader = append(funcsInHeader, fun)
		}
		r.Close()
	}

	// Extract Win32API's from {dllname}.def
	funcsInDLLs := getWinDllSymbols(reader)

	// funcs = funcsInDLLs ∩ funcsInHeader
	funcs := make(map[string][]Func)
	missingFuncs := []Func{}
	for _, dllFunc := range funcsInDLLs {
		found := false
		for _, headerFunc := range funcsInHeader {
			if dllFunc.Name != headerFunc.Name {
				continue
			}
			_, ok := funcs[dllFunc.Dll]
			if !ok {
				funcs[dllFunc.Dll] = make([]Func, 0)
			}
			funcs[dllFunc.Dll] = append(funcs[dllFunc.Dll], Func{dllFunc.Dll, dllFunc.Name, headerFunc.Args, headerFunc.Ret, dllFunc.IsUnicodeFunc})
			found = true
			break
		}
		if !found {
			missingFuncs = append(missingFuncs, *dllFunc)
		}
	}

	wineZipPath, err := downloadWineZip(wineVersion)
	if err != nil {
		panic(err)
	}

	// Search function definition from Wine's source files.
	funcs = digWineSourceFiles(wineZipPath, missingFuncs, funcs)

	// Remove functions which has va_list argument.
	tmp := make(map[string][]Func)
	for dll, dllFuncs := range funcs {
		tmp[dll] = make([]Func, 0)
		for _, fun := range dllFuncs {
			skip := false
			for _, arg := range fun.Args {
				name := arg.Type.Name
				if name == "va_list *" || name == "__ms_va_list" || name == "va_list" || name == "__ms_va_list *" || name == "..." {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			tmp[dll] = append(tmp[dll], fun)
		}
	}

	saveCache(tmp, cacheFilePath)

	return tmp
}

func createTypedefFile(referencedTypes *StringSet) error {
	types := NewStringSet()
	for _, t := range referencedTypes.Values() {
		if strings.HasPrefix(t, "*") {
			ty := NewType(t)
			t = ty.DereferencedGoName()
		}
		if strings.Contains(t, "/*") && strings.Contains(t, "*/") {
			continue
		}
		types.Put(t)
	}

	// collect all types used in struct.
	changed := true
	for changed {
		changed = false
		before := types.Size()
		for _, name := range types.Values() {
			if strings.Contains(name, "/*") && strings.Contains(name, "*/") {
				continue
			}
			templateFilePath := TemplateFilePath(name, "")
			templateFile, err := os.Open(templateFilePath)
			if err != nil {
				continue
			}
			s := bufio.NewScanner(templateFile)
			for s.Scan() {
				line := s.Text()
				if strings.HasPrefix(line, "package") || line == "" {
					continue
				}
				if strings.HasPrefix(line, "//") {
					if strings.HasPrefix(line, "//ref ") {
						tokens := strings.Split(line, "//ref ")
						t := strings.Trim(tokens[1], " ")
						types.Put(t)
					}
					continue
				}
			}
			templateFile.Close()
		}
		changed = before != types.Size()
	}

	typesFilePath := "types.go"
	typesFile, err := os.Create(typesFilePath)
	if err != nil {
		return err
	}
	defer typesFile.Close()
	fmt.Fprintf(typesFile, "package win\n")
	fmt.Fprintf(typesFile, "\n")

	importedPackages := NewStringSet()
	typeDefLines := []string{}

	reg := regexp.MustCompile(`\s*(LPCWSTR)\s*`)
	for _, name := range types.Values() {
		templateFilePath := TemplateFilePath(name, "")
		templateFile, err := os.Open(templateFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot find type definition file: %s\n", templateFilePath)
		}
		s := bufio.NewScanner(templateFile)
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, "package") || line == "" {
				continue
			}
			if strings.HasPrefix(line, "import ") {
				tokens := strings.Split(line, "import ")
				imp := strings.Trim(strings.Trim(tokens[1], " "), "\"")
				importedPackages.Put(imp)
				continue
			}
			if strings.HasPrefix(line, "//") {
				continue
			}
			if strings.Contains(line, "func") && !strings.HasPrefix(line, "func") {
				lineBytes := []byte(line)
				replaced := reg.ReplaceAllFunc(lineBytes, func(match []byte) []byte {
					return []byte(" string ")
				})
				line = string(replaced)
				line = strings.Replace(line, "  ", " ", -1)
				line = strings.Replace(line, " , ", ", ", -1)
				line = strings.Replace(line, " )", ")", -1)
				line = strings.Trim(line, " ")
			}
			typeDefLines = append(typeDefLines, line)
		}
		templateFile.Close()
	}

	importedPackages.Put("reflect")

	if importedPackages.Size() > 0 {
		fmt.Fprintf(typesFile, "import (\n")
		for _, imp := range importedPackages.Values() {
			fmt.Fprintf(typesFile, "\t\"%s\"\n", imp)
		}
		fmt.Fprintf(typesFile, ")\n")
		fmt.Fprintf(typesFile, "\n")
	}

	fmt.Fprintf(typesFile, "var (\n")
	fmt.Fprintf(typesFile, "\ttypes map[string]reflect.Type\n")
	fmt.Fprintf(typesFile, ")\n")

	fmt.Fprintf(typesFile, "func init() {\n")
	fmt.Fprintf(typesFile, "\ttypes = make(map[string]reflect.Type)\n")
	for _, t := range types.Values() {
		fmt.Fprintf(typesFile, "\ttypes[\"%s\"] = reflect.TypeOf((*%s)(nil)).Elem()\n", t, t)
	}
	fmt.Fprintf(typesFile, "}\n")
	fmt.Fprintf(typesFile, "\n")

	for _, line := range typeDefLines {
		fmt.Fprintf(typesFile, "%s\n", line)
	}

	fmt.Fprintf(typesFile, "func Typeof(name string) reflect.Type {\n")
	fmt.Fprintf(typesFile, "\treturn types[name]\n")
	fmt.Fprintf(typesFile, "}\n")
	fmt.Fprintf(typesFile, "\n")

	fmt.Fprintf(typesFile, "func TypeNames() []string {\n")
	fmt.Fprintf(typesFile, "\tret %s= make([]string, len(types))\n", ":")
	fmt.Fprintf(typesFile, "\ti %s= 0\n", ":")
	fmt.Fprintf(typesFile, "\tfor name, _ %s= range types {\n", ":")
	fmt.Fprintf(typesFile, "\t\tret[i] = name\n")
	fmt.Fprintf(typesFile, "\t\ti++\n")
	fmt.Fprintf(typesFile, "\t}\n")
	fmt.Fprintf(typesFile, "\treturn ret\n")
	fmt.Fprintf(typesFile, "}\n")
	fmt.Fprintf(typesFile, "\n")

	typesFile.Close()

	for _, arch := range []string{"386", "amd64"} {
		fp, err := os.Create(fmt.Sprintf("types_%s.go", arch))
		if err != nil {
			return err
		}
		localImports := NewStringSet()
		lines := []string{}
		fmt.Fprintf(fp, "package win\n")
		fmt.Fprintf(fp, "\n")
		for _, name := range types.Values() {
			templateFilePath := TemplateFilePath(name, arch)
			templateFile, err := os.Open(templateFilePath)
			if err != nil {
				continue
			}
			s := bufio.NewScanner(templateFile)
			for s.Scan() {
				line := s.Text()
				if strings.HasPrefix(line, "package") || line == "" {
					continue
				}
				if strings.HasPrefix(line, "//") {
					continue
				}
				if strings.HasPrefix(line, "import ") {
					tokens := strings.Split(line, "import ")
					imp := strings.Trim(strings.Trim(tokens[1], " "), "\"")
					localImports.Put(imp)
					continue
				}
				lines = append(lines, line)
			}
			templateFile.Close()
		}
		if localImports.Size() > 0 {
			fmt.Fprintf(fp, "import (\n")
			for _, imp := range localImports.Values() {
				fmt.Fprintf(fp, "\t\"%s\"\n", imp)
			}
			fmt.Fprintf(fp, ")\n")
			fmt.Fprintf(fp, "\n")
		}
		for _, line := range lines {
			fmt.Fprintf(fp, "%s\n", line)
		}
		fp.Close()
	}

	if err := gofmt(typesFilePath); err != nil {
		return err
	}

	return nil
}

func printFuncs(funcs []Func, dllName string) (requiredTypes *StringSet, printedFuncs []Func, e error) {
	requiredTypes = NewStringSet()
	goFilePath := fmt.Sprintf("%s.go", dllName)
	f, err := os.Create(goFilePath)
	if err != nil {
		return NewStringSet(), []Func{}, err
	}
	defer f.Close()

	// Headers
	header := []string{
		"// +build windows",
		"",
		"package win",
		"",
	}
	for _, h := range header {
		fmt.Fprintf(f, "%s\n", h)
	}

	libvarName := fmt.Sprintf("lib%s", strings.ToLower(dllName))

	exportedFuncs := []Func{}
	funcExportLines := []string{}
	importedPackages := NewStringSet()

	// Function definitions
	addEmptyLine := true
	for _, fun := range funcs {
		lines, types, imported, e := fun.ToString(fun.FuncPtrVarName(), 1)
		if len(lines) == 0 {
			if addEmptyLine {
				funcExportLines = append(funcExportLines, "")
			}
			addEmptyLine = false

			decl, _, _ := fun.Decl()
			funcExportLines = append(funcExportLines, fmt.Sprintf("// TODO: %s", e))
			funcExportLines = append(funcExportLines, fmt.Sprintf("// %s", decl))
			funcExportLines = append(funcExportLines, "")
			continue
		}

		addEmptyLine = true
		exportedFuncs = append(exportedFuncs, fun)
		requiredTypes.Merge(types)
		importedPackages.Merge(imported)

		funcExportLines = append(funcExportLines, "")
		for _, line := range lines {
			funcExportLines = append(funcExportLines, line)
		}
	}

	// import
	importedPackageNames := importedPackages.Values()
	if len(importedPackageNames) > 0 {
		fmt.Fprintf(f, "import (\n")
		for _, p := range importedPackageNames {
			fmt.Fprintf(f, "\t\"%s\"\n", p)
		}
		fmt.Fprintf(f, ")\n")
		fmt.Fprintf(f, "\n")
	}

	// DLL ptr, and Function ptrs
	fmt.Fprintf(f, "var (\n")
	fmt.Fprintf(f, "\t// Library\n")
	fmt.Fprintf(f, "\t%s uintptr\n", libvarName)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\t// Functions\n")
	for _, fun := range exportedFuncs {
		fmt.Fprintf(f, "\t%s uintptr\n", fun.FuncPtrVarName())
	}
	fmt.Fprintf(f, ")\n")

	// func init()
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\t// Library\n")
	fmt.Fprintf(f, "\t%s = doLoadLibrary(\"%s.dll\")\n", libvarName, dllName)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\t// Functions\n")
	for _, fun := range exportedFuncs {
		fmt.Fprintf(f, "\t%s = doGetProcAddress(%s, \"%s\")\n", fun.FuncPtrVarName(), libvarName, fun.Name)
	}
	fmt.Fprintf(f, "}\n")

	for _, line := range funcExportLines {
		fmt.Fprintf(f, "%s\n", line)
	}

	f.Close()
	gofmt(goFilePath)

	return requiredTypes, exportedFuncs, nil
}

func removeFileExt(filePath string) string {
	ext := filepath.Ext(filePath)
	if len(ext) == 0 {
		return filePath
	} else {
		return strings.TrimSuffix(filePath, ext)
	}
}

// Extract WINAPI function definitions from C header.
func extractWinAPIDefinition(r io.Reader) []Func {
	ret := []Func{}
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return ret
	}

	bytes = regexp.MustCompile(`//[^\n]*\n`).ReplaceAllFunc(bytes, func(match []byte) []byte {
		return []byte(" ")
	})
	bytes = regexp.MustCompile(`/\*.*?\*/`).ReplaceAllFunc(bytes, func(match []byte) []byte {
		return []byte(" ")
	})

	//WINIMPM HCERTSTORE WINAPI CertOpenStore (LPCSTR lpszStoreProvider, DWORD dwEncodingType, HCRYPTPROV_LEGACY hCryptProv, DWORD dwFlags, const void *pvPara);
	reg1 := regexp.MustCompile(`\s*(WINAPI|APIENTRY|WINGDIAPI|WINBASEAPI|WINUSERAPI)?\s*([A-Za-z0-9_]*)\s*(WINAPI|APIENTRY|WINGDIAPI|WINBASEAPI|WINUSERAPI)?\s*([A-Za-z0-9_]*)?\s*(WINAPI|APIENTRY|WINGDIAPI|WINBASEAPI|WINUSERAPI)?\s*\(([^)]*)\)\s*;`)

	match := reg1.FindAllSubmatch(bytes, -1)
	for i := 0; i < len(match); i++ {
		exportMarker1 := removeCRLF(string(match[i][1]))
		rettype1 := removeCRLF(string(match[i][2]))
		exportMarker2 := removeCRLF(string(match[i][3]))
		name := removeCRLF(string(match[i][4]))
		exportMarker3 := removeCRLF(string(match[i][5]))
		args := removeCRLF(string(match[i][6]))
		if exportMarker1 == "" && exportMarker2 == "" && exportMarker3 == "" {
			continue
		}
		ret = append(ret, Func{"", name, NewArgs(args), NewType(rettype1), false})
	}
	return ret
}

func removeCRLF(s string) string {
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.Replace(s, "\r", " ", -1)
	s = strings.Replace(s, "\t", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)
	s = strings.Trim(s, " ")
	return s
}

func digWineSourceFiles(wineZipPath string, missingFuncs []Func, funcs map[string][]Func) map[string][]Func {
	fp, err := os.Open(wineZipPath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	zipReader, err := zip.NewReader(fp, FileSize(wineZipPath))
	if err != nil {
		panic(err)
	}
	regCSourceFile := regexp.MustCompile(`^.*\.c$`)
	regLineComment := regexp.MustCompile(`//[^\n]*\n`)
	regMultiLineComment := regexp.MustCompile(`/\*.*?\*/`)
	regCFunctionDef := regexp.MustCompile(`\s*([A-Za-z0-9_]+)\s+WINAPI\s+([A-Za-z0-9_]+)\s*\(([^)]*)\)`)

	foundFuncDefs := []*Func{}
	for _, file := range zipReader.File {
		m := regCSourceFile.FindSubmatch([]byte(file.Name))
		if len(m) == 0 {
			continue
		}
		r, err := file.Open()
		if err != nil {
			continue
		}
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			continue
		}
		bytes = regLineComment.ReplaceAllFunc(bytes, func(match []byte) []byte {
			return []byte(" ")
		})
		bytes = regMultiLineComment.ReplaceAllFunc(bytes, func(match []byte) []byte {
			return []byte(" ")
		})
		match := regCFunctionDef.FindAllSubmatch(bytes, -1)
		for i := 0; i < len(match); i++ {
			rettype := string(match[i][1])
			name := string(match[i][2])
			args := removeCRLF(string(match[i][3]))
			f := NewFunc("", name, NewArgs(args), rettype, false)
			foundFuncDefs = append(foundFuncDefs, f)
		}
	}

	for _, missingFunc := range missingFuncs {
		missingFuncName := missingFunc.Name
		found := false
		for _, foundFuncDef := range foundFuncDefs {
			if missingFuncName == foundFuncDef.Name {
				_, ok := funcs[missingFunc.Dll]
				if !ok {
					funcs[missingFunc.Dll] = make([]Func, 0)
				}
				funcs[missingFunc.Dll] = append(funcs[missingFunc.Dll], Func{missingFunc.Dll, missingFunc.Name, foundFuncDef.Args, foundFuncDef.Ret, missingFunc.IsUnicodeFunc})
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Function definition not found: %s @ %s\n", missingFunc.Name, missingFunc.Dll)
		}
	}

	return funcs
}

func getWinDllSymbols(mingwZipReader *zip.Reader) []*Func {
	ret := []*Func{}

	dlls := []string{
		"advapi32.dll",
		"comctl32.dll",
		"comdlg32.dll",
		"gdi32.dll",
		"gdiplus.dll",
		"kernel32.dll",
		"ole32.dll",
		"oleaut32.dll",
		"opengl32.dll",
		"pdh.dll",
		"psapi.dll",
		"shell32.dll",
		"uxtheme.dll",
		"user32.dll",
		"winmm.dll",
		"ws2_32.dll",
		"shlwapi.dll",
		"imm32.dll",
		"version.dll",
		"crypt32.dll",
		"rpcrt4.dll",
		"iphlpapi.dll",
	}

	duplicatedMasterSymbols := map[string][]string{
		"advapi32.dll": {
			"CreateProcessAsUser",
			"RegCloseKey",
			"RegCopyTree",
			"RegCreateKeyEx",
			"RegDeleteKeyEx",
			"RegDeleteTree",
			"RegDeleteValue",
			"RegDisablePredefinedCacheEx",
			"RegEnumKeyEx",
			"RegEnumValue",
			"RegFlushKey",
			"RegGetKeySecurity",
			"RegGetValue",
			"RegLoadKey",
			"RegLoadMUIString",
			"RegNotifyChangeKeyValue",
			"RegOpenCurrentUser",
			"RegOpenKeyEx",
			"RegOpenUserClassesRoot",
			"RegQueryInfoKey",
			"RegQueryValueEx",
			"RegRestoreKey",
			"RegSaveKeyEx",
			"RegSetKeySecurity",
			"RegSetValueEx",
			"RegUnLoadKey",
			"SetThreadToken",
			"CryptGetDefaultProvider",
		},
		"shlwapi.dll": {
			"SHAllocShared",
			"SHFreeShared",
			"SHLockShared",
			"SHUnlockShared",
			"StrChrI",
			"StrChr",
			"StrCmpNI",
			"StrCmpN",
			"StrCpy",
			"StrRChr",
			"StrRStrI",
			"StrStrI",
			"StrCpyN",
			"StrStr",
			"StrRChrI",
		},
		"kernel32.dll": {
			"VerLanguageName",
			"DelayLoadFailureHook",
		},
		"oleaut32.dll": {
			"CreateErrorInfo",
			"SetErrorInfo",
			"GetErrorInfo",
		},
	}

	reg := regexp.MustCompile(`^([^/]+)/mingw-w64-crt/lib64/([^/]+)\.def$`)

	for _, dll := range dlls {
		var file *zip.File
		for _, f := range mingwZipReader.File {
			match := reg.FindSubmatch([]byte(f.FileHeader.Name))
			if len(match) < 3 {
				continue
			}
			name := string(match[2])
			if name+".dll" == dll {
				file = f
				break
			}
		}
		if file == nil {
			continue
		}
		symbols := NewStringSet()
		r, err := file.Open()
		if err != nil {
			continue
		}
		s := bufio.NewScanner(r)
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, ";") {
				continue
			}
			if strings.HasPrefix(line, "LIBRARY") {
				continue
			}
			if line == "EXPORTS" {
				continue
			}
			symbols.Put(line)
		}
		r.Close()

		for _, symbol := range symbols.Values() {
			if strings.HasPrefix(symbol, "_") {
				continue
			}
			found := false
			for masterDll, masterSymbols := range duplicatedMasterSymbols {
				if masterDll == dll {
					continue
				}
				for _, masterSymbol := range masterSymbols {
					if symbol == masterSymbol || symbol == masterSymbol+"A" || symbol == masterSymbol+"W" {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				ret = append(ret, NewFunc(dll, symbol, []Arg{}, "", false))
			}
		}
	}

	filtered := []*Func{}
	for i := 0; i < len(ret); i++ {
		fun := ret[i]
		switch fun.Name {
		case "DllUnregisterServer", "DllRegisterServer", "DllCanUnloadNow", "DllGetClassObject":
			// Special entry point for OLE
			continue
		case "DllInstall":
			// Special entry point for regsvr32
			continue
		case "DllGetVersion":
			continue
		}

		endsWithA := strings.HasSuffix(fun.Name, "A")
		endsWithW := strings.HasSuffix(fun.Name, "W")

		if !endsWithA && !endsWithW {
			found := false
			for j := 0; j < len(ret); j++ {
				if ret[j].Name == fun.Name+"W" || ret[j].Name == fun.Name+"A" {
					found = true
					break
				}
			}
			if !found {
				filtered = append(filtered, fun)
			}
		} else {
			searchBasename := fun.Name[0 : len(fun.Name)-1]
			var search string
			if endsWithW {
				search = searchBasename + "A"
			} else {
				search = searchBasename + "W"
			}
			found := false
			for j := 0; j < len(ret); j++ {
				if ret[j].Name == search {
					found = true
					break
				}
			}
			if found {
				if endsWithW {
					fun.IsUnicodeFunc = true
					filtered = append(filtered, fun)
				}
			} else {
				filtered = append(filtered, fun)
			}
		}
	}

	return filtered
}

func downloadMinGWZip(minGWVersion string) (string, error) {
	url := fmt.Sprintf("https://github.com/mirror/mingw-w64/archive/%s.zip", minGWVersion)
	zipPath := filepath.Join("internal", fmt.Sprintf("mingw-w64-%s.zip", minGWVersion))
	if !IsFileExist(zipPath) {
		if err := Wget(url, zipPath); err != nil {
			return "", err
		}
	}
	return zipPath, nil
}

func downloadWineZip(wineVersion string) (string, error) {
	url := fmt.Sprintf("https://github.com/wine-mirror/wine/archive/wine-%s.zip", wineVersion)
	zipPath := filepath.Join("internal", fmt.Sprintf("wine-%s.zip", wineVersion))
	if !IsFileExist(zipPath) {
		if err := Wget(url, zipPath); err != nil {
			return "", err
		}
	}
	return zipPath, nil
}

func win32headers(mingwZipReader *zip.Reader) []*zip.File {
	ret := []*zip.File{}

	//mingw-w64-5.0.0/mingw-w64-headers/include/bits.h
	reg := regexp.MustCompile(`^([^/]+)/mingw-w64-headers/include/(.+)\.h$`)
	for _, f := range mingwZipReader.File {
		match := reg.FindSubmatch([]byte(f.FileHeader.Name))
		if len(match) < 3 {
			continue
		}
		ret = append(ret, f)
	}

	return ret
}

func gofmt(file string) error {
	fin, err := os.Open(file)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(fin)
	if err != nil {
		fin.Close()
		return err
	}
	fin.Close()
	formatted, err := format.Source(bytes)
	if err != nil {
		return err
	}

	fout, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fout.Close()
	if err != nil {
		return err
	}
	fout.Write(formatted)
	return nil
}
