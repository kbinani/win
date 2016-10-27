package internal

import (
	"fmt"
	"strings"
)

func ReplaceString(s, old, new string) string {
	if strings.Contains(new, old) {
		panic(fmt.Errorf("new: '%s' contains old: '%s': this will cause infinite loop", new, old))
	}
	for strings.Contains(s, old) {
		s = strings.Replace(s, old, new, -1)
	}
	return s
}
