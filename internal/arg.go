package internal

import (
	"fmt"
	"strings"
)

type Arg struct {
	Type *Type
	Name string
}

func NewArg(arg string) Arg {
	arg = strings.Trim(arg, " ")
	arg = strings.Replace(arg, "*", " * ", -1)
	arg = strings.Trim(ReplaceString(arg, "  ", " "), " ")
	tokens := []string{}
	for _, token := range strings.Split(arg, " ") {
		if token == "IN" || token == "OUT" {
			continue
		}
		tokens = append(tokens, token)
	}
	if len(tokens) == 1 {
		return Arg{NewType(tokens[0]), ""}
	} else {
		var t, name string
		if tokens[len(tokens)-1] == "*" {
			name = ""
			t = strings.Join(tokens, " ")
		} else {
			name = tokens[len(tokens)-1]
			for i := 0; i < len(tokens)-1; i++ {
				t = t + tokens[i] + " "
			}
		}
		t = strings.Trim(t, " ")
		if strings.HasSuffix(name, "[]") {
			t = t + "*"
			name = strings.TrimSuffix(name, "[]")
		}
		return Arg{NewType(t), name}
	}
}

func NewArgs(args string) []Arg {
	ret := []Arg{}
	args = strings.Trim(args, " ")
	if args == "" {
		return ret
	}
	tokens := strings.Split(args, ",")
	for i, args := range tokens {
		arg := NewArg(args)
		if (arg.Type.Name == "VOID" || arg.Type.Name == "void") && len(tokens) == 1 {
			return ret
		}
		if arg.Name == "" {
			arg.Name = fmt.Sprintf("unnamed%d", i)
		}
		ret = append(ret, arg)
	}
	return ret
}

func (a Arg) GoName(index int) string {
	if a.Name == "" {
		return fmt.Sprintf("unnamed%d", index)
	}
	lname := strings.ToLower(a.Name)
	switch lname {
	case "func", "interface", "select", "defer", "go", "map", "chan", "package", "fallthrough", "range", "type", "import", "var", "len", "new":
		return "a" + strings.ToUpper(a.Name[0:1]) + a.Name[1:]
	default:
		return strings.ToLower(a.Name[0:1]) + a.Name[1:]
	}
}
