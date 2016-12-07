package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"unsafe"

	"bufio"

	"sync"

	win "github.com/kbinani/win"
	. "github.com/kbinani/win/internal"
)

var (
	blacklistSrc []string = []string{
		"MRUINFO",
		"NDR_SCONTEXT_",
		"INTERFACE_HANDLE",
		"REFCLSID",
		"REFGUID",
		"WSAEVENT",
		"REFIID",
		"PCIDLIST_ABSOLUTE", // alias of LPCITEMIDLIST
		"PCUITEMID_CHILD_ARRAY",
		"PCIDLIST_ABSOLUTE_ARRAY",
		"REFKNOWNFOLDERID",
		"PIDLIST_ABSOLUTE", // alias of LPITEMIDLIST
		"REFFMTID",
		"REFPROPVARIANT",

		"bool",
		"string",
		"int8",
		"uint8",
		"int16",
		"uint16",
		"int32",
		"uint32",
		"int64",
		"uint64",
		"float32",
		"float64",
		"byte",
		"int",
		"uintptr",
	}
	blacklist map[string]struct{}
)

func init() {
	blacklist = make(map[string]struct{})
	for _, v := range blacklistSrc {
		blacklist[v] = struct{}{}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	names := win.TypeNames()
	sort.Strings(names)

	for i := "A"[0]; i <= "Z"[0]; i++ {
		dir := filepath.Join("internal", "ctests", runtime.GOARCH, string(i))
		os.Mkdir(dir, 0777)
	}

	printedFiles := NewStringSet()

	var wg sync.WaitGroup
	m := new(sync.Mutex)
	for _, name := range names {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			filename, err := printTestCase(n)
			if err == nil && filename != "" {
				m.Lock()
				defer m.Unlock()
				printedFiles.Put(filename)
			}
		}(name)
	}
	wg.Wait()

	for i := "A"[0]; i <= "Z"[0]; i++ {
		dir := filepath.Join("internal", "ctests", runtime.GOARCH, string(i))
		d, err := os.Open(dir)
		if err != nil {
			panic(err)
		}
		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			panic(err)
		}
		for _, name := range names {
			n := filepath.Join(dir, name)
			if !printedFiles.Has(n) {
				os.Remove(n)
			}
		}
	}
}

func printTestCase(typename string) (string, error) {
	t := win.Typeof(typename)
	if t.Kind() == reflect.Func {
		return "", nil
	}
	_, ok := blacklist[typename]
	if ok {
		return "", fmt.Errorf("%s is blacklisted", typename)
	}

	switch typename {
	case "CRYPTOAPI_BLOB_":
		typename = "_CRYPTOAPI_BLOB"
	}

	lines := []string{}

	if runtime.GOARCH == "amd64" {
		lines = append(lines, fmt.Sprintf("#if defined(_WIN64)"))
	} else {
		lines = append(lines, fmt.Sprintf("#if !defined(_WIN64)"))
	}
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("#if defined(%s)", typename))
	lines = append(lines, fmt.Sprintf("\t#undef %s", typename))
	lines = append(lines, fmt.Sprintf("#endif"))
	lines = append(lines, fmt.Sprintf(""))

	draftTypeNames := []string{}
	draftTypeEnabler := []string{}

	lines = append(lines, fmt.Sprintf("namespace {"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("namespace size {"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("namespace of {"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("class impl {"))

	lines = append(lines, fmt.Sprintf("\tstruct order6 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order5 : public order6 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order4 : public order5 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order3 : public order4 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order2 : public order3 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order1 : public order2 {};"))
	lines = append(lines, fmt.Sprintf(""))

	enabler := fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%sW*>::value && std::is_pointer<::%sA*>::value>::type", typename, typename)
	lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
	lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order1 arg) { return sizeof(::%sW); }", typename, typename))
	lines = append(lines, fmt.Sprintf(""))
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%sW", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%s*>::value>::type", typename)
	lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
	lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order2 arg) { return sizeof(::%s); }", typename, typename))
	lines = append(lines, fmt.Sprintf(""))
	draftTypeNames = append(draftTypeNames, typename)
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%sW*>::value>::type", typename)
	lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
	lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order3 arg) { return sizeof(::%sW); }", typename, typename))
	lines = append(lines, fmt.Sprintf(""))
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%sW", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%s_W*>::value>::type", typename)
	lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
	lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order4 arg) { return sizeof(::%s_W); }", typename, typename))
	lines = append(lines, fmt.Sprintf(""))
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%s_W", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	lowerCamelTypename := strings.ToLower(typename[0:1]) + typename[1:]
	if lowerCamelTypename != typename {
		enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%s*>::value>::type", lowerCamelTypename)
		lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
		lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order5 arg) { return sizeof(::%s); }", typename, lowerCamelTypename))
		lines = append(lines, fmt.Sprintf(""))
		draftTypeNames = append(draftTypeNames, lowerCamelTypename)
		draftTypeEnabler = append(draftTypeEnabler, enabler)
	}
	if strings.Contains(typename, "_") {
		tokens := strings.Split(typename, "_")
		pre := strings.Join(tokens[0:len(tokens)-1], "_")
		versionedTypename := fmt.Sprintf("%sW_%s", pre, tokens[len(tokens)-1])

		enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<::%s*>::value>::type", versionedTypename)
		lines = append(lines, fmt.Sprintf("\ttemplate <%s>", enabler))
		lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order6 arg) { return sizeof(::%s); }", typename, versionedTypename))
		lines = append(lines, fmt.Sprintf(""))
		draftTypeNames = append(draftTypeNames, versionedTypename)
		draftTypeEnabler = append(draftTypeEnabler, enabler)
	}
	lines = append(lines, fmt.Sprintf("public:"))
	lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s() { return %s_impl(order1{}); }", typename, typename))

	lines = append(lines, fmt.Sprintf("};"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("static constexpr size_t %s = impl::%s();", typename, typename))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("} // namespace of"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("} // namespace size"))
	lines = append(lines, fmt.Sprintf(""))

	if t.Kind() == reflect.Struct {
		lines = append(lines, fmt.Sprintf("namespace offset {"))
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("namespace of {"))
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("namespace %s {", typename))
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("class impl {"))
		lines = append(lines, fmt.Sprintf("\ttemplate <typename>"))
		lines = append(lines, fmt.Sprintf("\tstruct enabler{"))
		lines = append(lines, fmt.Sprintf("\t\ttypedef bool type;"))
		lines = append(lines, fmt.Sprintf("\t};"))

		lines = append(lines, fmt.Sprintf(""))
		maxOrder := 2 * len(draftTypeNames)
		for order := maxOrder; order >= 1; order-- {
			if order == maxOrder {
				lines = append(lines, fmt.Sprintf("\tstruct order%d {};", order))
			} else {
				lines = append(lines, fmt.Sprintf("\tstruct order%d : public order%d {};", order, order+1))
			}
		}
		lines = append(lines, fmt.Sprintf(""))

		fields := getFields(t)
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
				continue
			}
			draftFieldNames := []string{}
			draftFieldNames = append(draftFieldNames, sf.Name)
			draftFieldNames = append(draftFieldNames, strings.ToLower(sf.Name[0:1])+sf.Name[1:])
			order := 1
			for j := 0; j < len(draftTypeNames); j++ {
				draftTypeName := draftTypeNames[j]
				typeEnabler := draftTypeEnabler[j]
				for _, draftFieldName := range draftFieldNames {
					lines = append(lines, fmt.Sprintf("\ttemplate <%s, typename = enabler<decltype(::%s::%s)>::type>", typeEnabler, draftTypeName, draftFieldName))
					lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s_impl(order%d arg0) { return offsetof(::%s, %s); }", sf.Name, order, draftTypeName, draftFieldName))
					lines = append(lines, fmt.Sprintf(""))
					order++
				}
			}
		}
		lines = append(lines, fmt.Sprintf("public:"))
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
				continue
			}
			lines = append(lines, fmt.Sprintf("\tstatic constexpr size_t %s() { return %s_impl(order1{}); }", sf.Name, sf.Name))
		}

		lines = append(lines, fmt.Sprintf("};"))
		lines = append(lines, fmt.Sprintf(""))
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			lines = append(lines, fmt.Sprintf("static constexpr size_t %s = impl::%s();", sf.Name, sf.Name))
		}
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("} // namespace %s", typename))
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("} // namespace of"))
		lines = append(lines, fmt.Sprintf(""))
		lines = append(lines, fmt.Sprintf("} // namespace offset"))
		lines = append(lines, fmt.Sprintf(""))
	}

	lines = append(lines, fmt.Sprintf("} // anonymouse namespace"))
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("TEST(%s, size) {", typename))
	lines = append(lines, fmt.Sprintf("\tEXPECT_EQ(size::of::%s, %d);", typename, t.Size()))
	lines = append(lines, fmt.Sprintf("}"))
	lines = append(lines, fmt.Sprintf(""))

	if t.Kind() == reflect.Struct {
		fields := getFields(t)
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
				continue
			}
			lines = append(lines, fmt.Sprintf("TEST(%s, offsetof_%s) {", typename, sf.Name))
			lines = append(lines, fmt.Sprintf("\tEXPECT_EQ(offset::of::%s::%s, %d);", typename, sf.Name, sf.Offset))
			lines = append(lines, fmt.Sprintf("}"))
			lines = append(lines, fmt.Sprintf(""))
		}
	}

	// lines = append(lines, fmt.Sprintf("#endif // !defined(%s)", typename))
	if runtime.GOARCH == "amd64" {
		lines = append(lines, fmt.Sprintf("#endif // defined(_WIN64)"))
	} else {
		lines = append(lines, fmt.Sprintf("#endif // !defined(_WIN64)"))
	}

	pre := typename[0:1]
	filename := filepath.Join("internal", "ctests", runtime.GOARCH, pre, fmt.Sprintf("%s_%s_test.cpp", TemplateFileName(typename), runtime.GOARCH))
	if isFileSame(lines, filename) {
		return filename, nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	for _, line := range lines {
		fmt.Fprintf(f, "%s\n", line)
	}

	return filename, nil
}

func isFileSame(expectedLines []string, actualFileName string) bool {
	f, err := os.Open(actualFileName)
	if err != nil {
		return false
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() && i < len(expectedLines) {
		line := s.Text()
		if line != expectedLines[i] {
			return false
		}
		i++
	}
	return i == len(expectedLines)
}

type tagField struct {
	Name   string
	Offset uintptr
}

func getFields(t reflect.Type) []tagField {
	ret := []tagField{}
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
			continue
		}
		ret = append(ret, tagField{sf.Name, sf.Offset})
	}

	ptrType := reflect.PtrTo(t)
	v := reflect.New(t)
	base := v.Pointer()
	for i := 0; i < ptrType.NumMethod(); i++ {
		m := ptrType.Method(i)
		for j := 0; j < m.Type.NumOut(); j++ {
			o := m.Type.Out(j)
			if o.Kind() != reflect.Ptr {
				continue
			}
			if strings.HasPrefix(m.Name, "Get") {
				continue
			}
			r := m.Func.Call([]reflect.Value{v})
			offset := r[0].Pointer() - base
			ret = append(ret, tagField{m.Name, offset})
		}
	}

	getters := NewStringSet()
	setters := NewStringSet()
	for i := 0; i < ptrType.NumMethod(); i++ {
		m := ptrType.Method(i)
		if strings.HasPrefix(m.Name, "Get") {
			if m.Type.NumOut() != 1 {
				continue
			}
			if m.Type.NumIn() != 1 {
				continue
			}
			getters.Put(strings.TrimPrefix(m.Name, "Get"))
		} else if strings.HasPrefix(m.Name, "Set") {
			if m.Type.NumOut() != 0 {
				continue
			}
			if m.Type.NumIn() != 2 {
				continue
			}
			setters.Put(strings.TrimPrefix(m.Name, "Set"))
		}
	}

	props := NewStringSet()
	for _, g := range getters.Values() {
		if !setters.Has(g) {
			continue
		}
		getter, _ := ptrType.MethodByName(fmt.Sprintf("Get%s", g))
		setter, _ := ptrType.MethodByName(fmt.Sprintf("Set%s", g))
		gtype := getter.Type.Out(0)
		stype := setter.Type.In(1)
		if gtype.String() == stype.String() {
			props.Put(g)
		}
	}
	for _, p := range props.Values() {
		for i := 0; i < int(t.Size()); i++ {
			*(*byte)(unsafe.Pointer(base + uintptr(i))) = 0
		}

		setter, _ := ptrType.MethodByName(fmt.Sprintf("Set%s", p))
		ptype := setter.Type.In(1)
		pvalue := reflect.New(ptype)
		for i := 0; i < int(ptype.Size()); i++ {
			*(*byte)(unsafe.Pointer(pvalue.Pointer() + uintptr(i))) = 0xFF
		}
		setter.Func.Call([]reflect.Value{v, reflect.Indirect(pvalue)})

		offset := 0
		for i := 0; i < int(t.Size()); i++ {
			b := *(*byte)(unsafe.Pointer(base + uintptr(i)))
			if b != 0 {
				offset = i
				break
			}
		}
		ret = append(ret, tagField{p, uintptr(offset)})
	}

	return ret
}
