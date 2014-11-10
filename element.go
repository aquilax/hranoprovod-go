package main

import (
	"sort"
)

type Element struct {
	name string
	val  float32
}

type Elements []*Element

func NewElement(name string, val float32) *Element {
	return &Element{name, val}
}

func NewElements() *Elements {
	return &Elements{}
}

func (el *Elements) Add(name string, val float32) {
	*el = append(*el, NewElement(name, val))
}

func (el *Elements) Index(name string) (int, bool) {
	for n, e := range *el {
		if e.name == name {
			return n, true
		}
	}
	return 0, false
}

func (els *Elements) SumMerge(left *Elements, coef float32) {
	for _, v := range *left {
		if ndx, exists := (*els).Index(v.name); exists {
			(*els)[ndx].val += v.val * coef
		} else {
			(*els).Add(v.name, v.val*coef)
		}
	}
}

func (e Elements) Len() int {
	return len(e)
}
func (e Elements) Less(i, j int) bool {
	return e[i].name < e[j].name
}
func (e Elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (el Elements) Sort() {
	sort.Sort(el)
}
