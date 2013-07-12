package main

import (
	"sort"
)

type Element struct {
	name string
	val  float32
}

type Elements []Element

func (el *Elements) Index(name string) (int, bool) {
	for n, e := range *el {
		if e.name == name {
			return n, true
		}
	}
	return 0, false
}

func (el *Elements) Add(name string, val float32) {
	var e Element
	e.name = name
	e.val = val
	*el = append(*el, e)
}

func (s Elements) Len() int {
	return len(s)
}
func (s Elements) Less(i, j int) bool {
	return s[i].name < s[j].name
}
func (s Elements) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (el Elements) Sort() {
	sort.Sort(el)
}
