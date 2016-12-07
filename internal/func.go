package internal

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Func struct {
	Dll           string
	Name          string
	Args          []Arg
	Ret           *Type
	IsUnicodeFunc bool
}

var (
	regFuncDecl *regexp.Regexp
	regComment  *regexp.Regexp
)

func init() {
	regFuncDecl = regexp.MustCompile(`^func\s*\((.*)\)\s*(.*)?$`)
	regComment = regexp.MustCompile(`/\*(.*?)\*/`)
}

func NewFunc(dll string, name string, args []Arg, ret string, isUnicodeFunc bool) *Func {
	r := new(Func)
	r.Dll = dll
	r.Name = name
	r.Args = args
	r.Ret = NewType(ret)
	r.IsUnicodeFunc = isUnicodeFunc
	return r
}

func NewFuncFromDecl(decl string) *Func {
	this := new(Func)
	bytes := regComment.ReplaceAllFunc([]byte(decl), func(match []byte) []byte {
		return []byte(" ")
	})
	decl = string(bytes)
	decl = ReplaceString(decl, "  ", " ")
	m := regFuncDecl.FindSubmatch([]byte(decl))
	argsStr := strings.Trim(string(m[1]), " ")
	retStr := string(m[2])
	this.Ret = NewType(retStr)
	if len(argsStr) > 0 {
		tokens := strings.Split(argsStr, ",")
		for _, token := range tokens {
			token = strings.Replace(token, "*", " * ", -1)
			token = ReplaceString(token, "  ", " ")
			token = strings.Trim(token, " ")
			argTokens := strings.Split(token, " ")
			argTokens[0], argTokens[len(argTokens)-1] = argTokens[len(argTokens)-1], argTokens[0]
			this.Args = append(this.Args, NewArg(strings.Join(argTokens, " ")))
		}
	}
	return this
}

func (this Func) GoName() string {
	if this.Name == "" {
		return ""
	}
	if this.IsUnicodeFunc {
		return strings.ToUpper(this.Name[0:1]) + this.Name[1:len(this.Name)-1]
	}
	return strings.ToUpper(this.Name[0:1]) + this.Name[1:]
}

func (fun Func) unpackArgs() []Arg {
	ret := []Arg{}
	for i, arg := range fun.Args {
		name := arg.GoName(i)
		typeName := arg.Type.Name
		goTypeName := arg.Type.GoName()
		if typeName == "POINT" {
			ret = append(ret, Arg{NewType("int32"), fmt.Sprintf("%s.X", name)})
			ret = append(ret, Arg{NewType("int32"), fmt.Sprintf("%s.Y", name)})
		} else if typeName == "SIZEL" || typeName == "SIZE" {
			ret = append(ret, Arg{NewType("int32"), fmt.Sprintf("%s.Cx", name)})
			ret = append(ret, Arg{NewType("int32"), fmt.Sprintf("%s.Cy", name)})
		} else if typeName == "COORD" {
			ret = append(ret, Arg{NewType("uintptr"), fmt.Sprintf("getUintptrFromCOORD(%s)", name)})
		} else if typeName == "BLENDFUNCTION" {
			ret = append(ret, Arg{NewType("uintptr"), fmt.Sprintf("getUintptrFromBLENDFUNCTION(%s)", name)})
		} else if typeName == "GUID" || typeName == "CLSID" {
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("%s.Data1", name)})
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("(uint32(%s.Data2) << 16) | uint32(%s.Data3)", name, name)})
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("(uint32(%s.Data4[0]) << 24) | (uint32(%s.Data4[1] << 16)) | (uint32(%s.Data4[2] << 8)) | uint32(%s.Data4[3])", name, name, name, name)})
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("(uint32(%s.Data4[4]) << 24) | (uint32(%s.Data4[5] << 16)) | (uint32(%s.Data4[6] << 8)) | uint32(%s.Data4[7])", name, name, name, name)})
		} else if goTypeName == "CY" || goTypeName == "ULARGE_INTEGER" {
			// struct with 8 bytes size
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("*(*uint32)(unsafe.Pointer(&%s))", name)})
			ret = append(ret, Arg{NewType("uint32"), fmt.Sprintf("*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&%s)) + uintptr(4)))", name)})
		} else {
			ret = append(ret, arg)
		}
	}
	return ret
}

func (fun Func) Decl() (declLine string, referencedTypes *StringSet, e error) {
	decl := fmt.Sprintf("func %s(", fun.GoName())
	args := []string{}
	types := NewStringSet()

	unknownTypes := NewStringSet()

	for i, arg := range fun.Args {
		goType := arg.Type.GoName()
		if goType == "" {
			unknownTypes.Put(strings.Replace(arg.Type.Name, "  ", " ", -1))
			goType = arg.Type.Name
		} else {
			types.Put(goType)
		}
		if goType == "LPCWSTR" {
			goType = "string"
		}
		constComment := " "
		if arg.Type.IsConst() {
			constComment = " /*const*/ "
		}
		args = append(args, fmt.Sprintf("%s%s%s", arg.GoName(i), constComment, goType))
	}
	decl = decl + fmt.Sprintf("%s)", strings.Join(args, ", "))

	retGoType := fun.Ret.GoName()
	retIsVoid := retGoType == "void"
	if retGoType == "" && !retIsVoid {
		unknownTypes.Put(fun.Ret.Name)
		retGoType = fun.Ret.Name
	}
	if retGoType == "LPCWSTR" {
		retGoType = "string"
	}
	if !retIsVoid {
		decl = decl + fmt.Sprintf(" %s", retGoType)
		types.Put(retGoType)
	}

	var err error
	if unknownTypes.Size() > 0 {
		err = fmt.Errorf("Unknown type(s): %s", strings.Join(unknownTypes.Values(), ", "))
	}
	return decl, types, err
}

func (fun Func) ToString(funcPtrVarName string, indentLevel int) (lines []string, referencedTypes *StringSet, importedPackages *StringSet, e error) {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	ret := []string{}
	types := NewStringSet()
	imported := NewStringSet()

	decl, declTypes, err := fun.Decl()
	if err != nil {
		return []string{}, NewStringSet(), NewStringSet(), err
	}
	types.Merge(declTypes)

	templateFile := filepath.Join("internal", "funcs", fmt.Sprintf("%s.go", fun.GoName()))
	if IsFileExist(templateFile) {
		f, e := os.Open(templateFile)
		if e != nil {
			return []string{}, NewStringSet(), NewStringSet(), e
		}
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, "package") || line == "" {
				continue
			}
			ret = append(ret, line)
		}
		return ret, types, imported, nil
	}

	retGoType := fun.Ret.GoName()
	retIsVoid := retGoType == "void"

	decl = decl + " {"
	ret = append(ret, decl)

	unpackedArgs := fun.unpackArgs()
	argTypes := []string{}
	for _, arg := range unpackedArgs {
		goType := arg.Type.GoName()
		if goType == "" {
			return []string{}, NewStringSet(), NewStringSet(), fmt.Errorf("Cannot get Go typename: %s", arg.Type.Name)
		}
		argTypes = append(argTypes, goType)
	}

	// convert string to []uint16
	for i, arg := range unpackedArgs {
		if arg.Type.GoName() != "LPCWSTR" {
			continue
		}
		name := arg.GoName(i)
		ret = append(ret, fmt.Sprintf("%s%sStr %s= unicode16FromString(%s)", indent, name, ":", name))
	}

	// create callbacks
	for i, arg := range unpackedArgs {
		funcDecl := arg.Type.FuncDecl()
		if funcDecl == "" {
			continue
		}
		argName := arg.GoName(i)

		match := regFuncDecl.FindSubmatch([]byte(funcDecl))
		callbackArgStr := string(match[1])
		retType := string(match[2])
		callbackArgs := strings.Split(callbackArgStr, ", ")

		if retType == "uintptr" || retType == "" {
			/*
				argNameCallback := syscall.NewCallback(argName)
			*/
			ret = append(ret, fmt.Sprintf("%s%sCallback %s= syscall.NewCallback(%s)", indent, argName, ":", argName))
			imported.Put("syscall")
		} else {
			/*
				argNameCallback := syscall.NewCallback(func(argName1 argType1, argName2 argType2, ...) {
					ret := argName(argName1, argName2, ...)
					return uintptr(ret)
				})
			*/
			typeConversionLines := []string{}
			declArgs := []string{}
			callArgs := []string{}

			if callbackArgStr != "" {
				for _, a := range callbackArgs {
					nameAndType := strings.Split(a, " ")
					callArgName := nameAndType[0]
					callArgType := strings.Join(nameAndType[1:], " ")
					if callArgType == "LPCWSTR" {
						declArgs = append(declArgs, fmt.Sprintf("%sRawArg /*const*/ *uint16", callArgName))
						typeConversionLines = append(typeConversionLines, fmt.Sprintf("%s\t%s %s= stringFromUnicode16(%sRawArg)", indent, callArgName, ":", callArgName))
						callArgs = append(callArgs, callArgName)
					} else {
						declArgs = append(declArgs, fmt.Sprintf("%sRawArg %s", callArgName, callArgType))
						callArgs = append(callArgs, fmt.Sprintf("%sRawArg", callArgName))
					}
				}
			}

			ret = append(ret, fmt.Sprintf("%s%sCallback %s= syscall.NewCallback(func(%s) uintptr {", indent, argName, ":", strings.Join(declArgs, ", ")))
			imported.Put("syscall")
			for _, line := range typeConversionLines {
				ret = append(ret, line)
			}
			ret = append(ret, fmt.Sprintf("%s\tret %s= %s(%s)", indent, ":", argName, strings.Join(callArgs, ", ")))
			retGoType := NewType(retType)
			if retGoType.GoName() == "uintptr" {
				ret = append(ret, fmt.Sprintf("%s\treturn ret", indent))
			} else if retGoType.IsPtr() {
				ret = append(ret, fmt.Sprintf("%s\treturn uintptr(unsafe.Pointer(ret))", indent))
			} else {
				ret = append(ret, fmt.Sprintf("%s\treturn uintptr(ret)", indent))
			}
			ret = append(ret, fmt.Sprintf("%s})", indent))
		}

		if callbackArgStr != "" {
			for _, a := range callbackArgs {
				nameAndType := strings.Split(a, " ")
				typeName := nameAndType[1]
				types.Put(typeName)
			}
		}
	}

	syscallVariant := int(math.Max(1.0, math.Ceil(float64(len(unpackedArgs))/3.0))) * 3
	if syscallVariant > 15 {
		// syscall.Syscall18 (or over) is not available
		return []string{}, NewStringSet(), NewStringSet(), fmt.Errorf("Too many syscall arguments: %d", syscallVariant)
	}
	syscallName := fmt.Sprintf("syscall%d", syscallVariant)
	retVarName := fmt.Sprintf("ret%d", indentLevel)
	retAssign := fmt.Sprintf("%s %s= ", retVarName, ":")
	if retIsVoid {
		retAssign = ""
	}
	ret = append(ret, fmt.Sprintf("\t%s%s(%s, %d,", retAssign, syscallName, funcPtrVarName, len(unpackedArgs)))
	for i := 0; i < syscallVariant; i++ {
		line := ""
		if i < len(unpackedArgs) {
			arg := unpackedArgs[i]
			argName := arg.GoName(i)
			if argTypes[i] == "bool" {
				line = fmt.Sprintf("%s\tgetUintptrFromBool(%s)", indent, argName)
			} else if argTypes[i] == "uintptr" {
				line = fmt.Sprintf("%s\t%s", indent, argName)
			} else if arg.Type.GoName() == "LPCWSTR" {
				line = fmt.Sprintf("%s\tuintptr(unsafe.Pointer(&%sStr[0]))", indent, argName)
				imported.Put("unsafe")
			} else if arg.Type.IsPtr() {
				line = fmt.Sprintf("%s\tuintptr(unsafe.Pointer(%s))", indent, argName)
				imported.Put("unsafe")
			} else if arg.Type.FuncDecl() != "" {
				line = fmt.Sprintf("%s\t%sCallback", indent, argName)
				imported.Put("syscall")
			} else {
				line = fmt.Sprintf("%s\tuintptr(%s)", indent, argName)
			}
		} else {
			line = fmt.Sprintf("%s\t0", indent)
		}
		if i < syscallVariant-1 {
			line = line + ","
		} else {
			line = line + ")"
		}
		ret = append(ret, line)
	}
	if !retIsVoid {
		if retGoType == "bool" {
			ret = append(ret, fmt.Sprintf("%sreturn %s != 0", indent, retVarName))
		} else if retGoType == "COORD" {
			ret = append(ret, fmt.Sprintf("%sreturn getCOORDFromUintptr(%s)", indent, retVarName))
		} else if fun.Ret.GoName() == "LPCWSTR" {
			ret = append(ret, fmt.Sprintf("%sreturn stringFromUnicode16((*uint16)(unsafe.Pointer(%s)))", indent, retVarName))
		} else if fun.Ret.IsPtr() {
			ret = append(ret, fmt.Sprintf("%sreturn (%s)(unsafe.Pointer(%s))", indent, retGoType, retVarName))
			imported.Put("unsafe")
		} else if fun.Ret.FuncDecl() != "" {
			decl := fun.Ret.FuncDecl()
			fn := NewFuncFromDecl(decl)
			lines, fnTypes, fnImports, e := fn.ToString(retVarName, indentLevel+1)
			types.Merge(fnTypes)
			imported.Merge(fnImports)
			if e != nil {
				return []string{}, NewStringSet(), NewStringSet(), e
			}
			for i, line := range lines {
				if i == 0 {
					ret = append(ret, fmt.Sprintf("%sreturn %s", indent, line))
				} else {
					ret = append(ret, fmt.Sprintf("%s%s", indent, line))
				}
			}
		} else {
			ret = append(ret, fmt.Sprintf("%sreturn %s(%s)", indent, retGoType, retVarName))
		}
	}
	ret = append(ret, "}")

	return ret, types, imported, nil
}

func (this Func) FuncPtrVarName() string {
	if this.Name == "" {
		return ""
	}
	if this.IsUnicodeFunc {
		return strings.ToLower(this.Name[0:1]) + this.Name[1:len(this.Name)-1]
	}
	return strings.ToLower(this.Name[0:1]) + this.Name[1:]
}
