package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"

	win "github.com/kbinani/win"
)

var (
	blacklistSrc []string = []string{
		"MRUINFO",
		"NDR_SCONTEXT_",
		"INTERFACE_HANDLE",

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
		os.RemoveAll(dir)
		os.Mkdir(dir, 0777)
	}

	for _, name := range names {
		printTestCase(name)
	}
}

func printTestCase(typename string) error {
	t := win.Typeof(typename)
	if t.Kind() == reflect.Func {
		return nil
	}
	_, ok := blacklist[typename]
	if ok {
		return fmt.Errorf("%s is blacklisted", typename)
	}

	switch typename {
	case "PROPSHEETHEADER_V2":
		typename = "PROPSHEETHEADER"
	case "CRYPTOAPI_BLOB_":
		typename = "_CRYPTOAPI_BLOB"
	}

	pre := typename[0:1]
	filename := filepath.Join("internal", "ctests", runtime.GOARCH, pre, fmt.Sprintf("%s_%s_test.cpp", typename, runtime.GOARCH))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if runtime.GOARCH == "amd64" {
		fmt.Fprintf(f, "#if defined(_WIN64)\n")
	} else {
		fmt.Fprintf(f, "#if !defined(_WIN64)\n")
	}
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "#if !defined(%s)\n", typename)
	fmt.Fprintf(f, "\n")

	draftTypeNames := []string{}
	draftTypeEnabler := []string{}

	fmt.Fprintf(f, "namespace {\n")
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "namespace size {\n")
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "namespace of {\n")
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "class impl {\n")

	fmt.Fprintf(f, "\tstruct order6 {};\n")
	fmt.Fprintf(f, "\tstruct order5 : public order6 {};\n")
	fmt.Fprintf(f, "\tstruct order4 : public order5 {};\n")
	fmt.Fprintf(f, "\tstruct order3 : public order4 {};\n")
	fmt.Fprintf(f, "\tstruct order2 : public order3 {};\n")
	fmt.Fprintf(f, "\tstruct order1 : public order2 {};\n")
	fmt.Fprintf(f, "\n")

	enabler := fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%sW)>::value && std::is_pointer<decltype(new ::%sA)>::value>::type", typename, typename)
	fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
	fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order1 arg) { return sizeof(::%sW); }\n", typename, typename)
	fmt.Fprintf(f, "\n")
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%sW", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%s)>::value>::type", typename)
	fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
	fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order2 arg) { return sizeof(::%s); }\n", typename, typename)
	fmt.Fprintf(f, "\n")
	draftTypeNames = append(draftTypeNames, typename)
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%sW)>::value>::type", typename)
	fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
	fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order3 arg) { return sizeof(::%sW); }\n", typename, typename)
	fmt.Fprintf(f, "\n")
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%sW", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%s_W)>::value>::type", typename)
	fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
	fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order4 arg) { return sizeof(::%s_W); }\n", typename, typename)
	fmt.Fprintf(f, "\n")
	draftTypeNames = append(draftTypeNames, fmt.Sprintf("%s_W", typename))
	draftTypeEnabler = append(draftTypeEnabler, enabler)

	lowerCamelTypename := strings.ToLower(typename[0:1]) + typename[1:]
	if lowerCamelTypename != typename {
		enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%s)>::value>::type", lowerCamelTypename)
		fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
		fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order5 arg) { return sizeof(::%s); }\n", typename, lowerCamelTypename)
		fmt.Fprintf(f, "\n")
		draftTypeNames = append(draftTypeNames, lowerCamelTypename)
		draftTypeEnabler = append(draftTypeEnabler, enabler)
	}
	if strings.Contains(typename, "_") {
		tokens := strings.Split(typename, "_")
		pre := strings.Join(tokens[0:len(tokens)-1], "_")
		versionedTypename := fmt.Sprintf("%sW_%s", pre, tokens[len(tokens)-1])

		enabler = fmt.Sprintf("typename = std::enable_if<std::is_pointer<decltype(new ::%s)>::value>::type", versionedTypename)
		fmt.Fprintf(f, "\ttemplate <%s>\n", enabler)
		fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order6 arg) { return sizeof(::%s); }\n", typename, versionedTypename)
		fmt.Fprintf(f, "\n")
		draftTypeNames = append(draftTypeNames, versionedTypename)
		draftTypeEnabler = append(draftTypeEnabler, enabler)
	}
	fmt.Fprintf(f, "public:\n")
	fmt.Fprintf(f, "\tstatic constexpr size_t %s() { return %s_impl(order1{}); }\n", typename, typename)

	fmt.Fprintf(f, "};\n")
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "static constexpr size_t %s = impl::%s();\n", typename, typename)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "} // namespace of\n")
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "} // namespace size\n")
	fmt.Fprintf(f, "\n")

	if t.Kind() == reflect.Struct {
		fmt.Fprintf(f, "namespace offset {\n")
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "namespace of {\n")
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "namespace %s {\n", typename)
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "class impl {\n")
		fmt.Fprintf(f, "\ttemplate <typename>\n")
		fmt.Fprintf(f, "\tstruct enabler{\n")
		fmt.Fprintf(f, "\t\ttypedef bool type;\n")
		fmt.Fprintf(f, "\t};\n")

		fmt.Fprintf(f, "\n")
		maxOrder := 2 * len(draftTypeNames)
		for order := maxOrder; order >= 1; order-- {
			if order == maxOrder {
				fmt.Fprintf(f, "\tstruct order%d {};\n", order)
			} else {
				fmt.Fprintf(f, "\tstruct order%d : public order%d {};\n", order, order+1)
			}
		}
		fmt.Fprintf(f, "\n")

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
					fmt.Fprintf(f, "\ttemplate <%s, typename = enabler<decltype(::%s::%s)>::type>\n", typeEnabler, draftTypeName, draftFieldName)
					fmt.Fprintf(f, "\tstatic constexpr size_t %s_impl(order%d arg0) { return offsetof(::%s, %s); }\n", sf.Name, order, draftTypeName, draftFieldName)
					fmt.Fprintf(f, "\n")
					order++
				}
			}
		}
		fmt.Fprintf(f, "public:\n")
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
				continue
			}
			fmt.Fprintf(f, "\tstatic constexpr size_t %s() { return %s_impl(order1{}); }\n", sf.Name, sf.Name)
		}

		fmt.Fprintf(f, "};\n")
		fmt.Fprintf(f, "\n")
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			fmt.Fprintf(f, "static constexpr size_t %s = impl::%s();\n", sf.Name, sf.Name)
		}
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "} // namespace %s\n", typename)
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "} // namespace of\n")
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "} // namespace offset\n")
		fmt.Fprintf(f, "\n")
	}

	fmt.Fprintf(f, "} // anonymouse namespace\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "TEST(%s, size) {\n", typename)
	fmt.Fprintf(f, "\tEXPECT_EQ(size::of::%s, %d);\n", typename, t.Size())
	fmt.Fprintf(f, "}\n")
	fmt.Fprintf(f, "\n")

	if t.Kind() == reflect.Struct {
		fields := getFields(t)
		for i := 0; i < len(fields); i++ {
			sf := fields[i]
			if strings.ToLower(sf.Name[0:1]) == sf.Name[0:1] {
				continue
			}
			fmt.Fprintf(f, "TEST(%s, offsetof_%s) {\n", typename, sf.Name)
			fmt.Fprintf(f, "\tEXPECT_EQ(offset::of::%s::%s, %d);\n", typename, sf.Name, sf.Offset)
			fmt.Fprintf(f, "}\n")
			fmt.Fprintf(f, "\n")
		}
	}

	fmt.Fprintf(f, "#endif // !defined(%s)\n", typename)
	if runtime.GOARCH == "amd64" {
		fmt.Fprintf(f, "#endif // defined(_WIN64)\n")
	} else {
		fmt.Fprintf(f, "#endif // !defined(_WIN64)\n")
	}
	return nil
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
			r := m.Func.Call([]reflect.Value{v})
			offset := r[0].Pointer() - base
			ret = append(ret, tagField{m.Name, offset})
		}
	}
	return ret
}
