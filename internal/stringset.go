package internal

import (
	"sort"
)

type StringSet struct {
	values map[string]struct{}
}

func NewStringSet() *StringSet {
	o := new(StringSet)
	o.values = make(map[string]struct{})
	return o
}

func (o *StringSet) Put(s string) {
	o.values[s] = struct{}{}
}

func (o *StringSet) Values() []string {
	ret := []string{}
	for v := range o.values {
		ret = append(ret, v)
	}
	sort.Strings(ret)
	return ret
}

func (o *StringSet) Merge(other *StringSet) {
	for v := range other.values {
		o.values[v] = struct{}{}
	}
}

func (o *StringSet) Size() int {
	return len(o.values)
}

func (o *StringSet) Has(s string) bool {
	_, ok := o.values[s]
	return ok
}
