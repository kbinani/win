package internal

import (
	"reflect"
	"sort"
)

func SortedStringKeys(m interface{}) []string {
	value := reflect.ValueOf(m)
	if value.Kind() != reflect.Map {
		return []string{}
	}
	keys := value.MapKeys()
	ret := []string{}
	for _, key := range keys {
		if key.Kind() != reflect.String {
			continue
		}
		ret = append(ret, key.String())
	}
	sort.Strings(ret)
	return ret
}
