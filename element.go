package main

import (
	"sort"
)

// Element contains single element data
type Element struct {
	name string
	val  float32
}

// Elements contains array of elements
type Elements []*Element

// NewElement creates new element
func NewElement(name string, val float32) *Element {
	return &Element{name, val}
}

// NewElements creates new element list
func NewElements() *Elements {
	return &Elements{}
}

func (el *Elements) add(name string, val float32) {
	*el = append(*el, NewElement(name, val))
}

func (el *Elements) index(name string) (int, bool) {
	for n, e := range *el {
		if e.name == name {
			return n, true
		}
	}
	return 0, false
}

func (el *Elements) sumMerge(left *Elements, coef float32) {
	for _, v := range *left {
		if ndx, exists := (*el).index(v.name); exists {
			(*el)[ndx].val += v.val * coef
		} else {
			(*el).add(v.name, v.val*coef)
		}
	}
}

// Len returns the length of the element list
func (el Elements) Len() int {
	return len(el)
}

// Less compares two elements
func (el Elements) Less(i, j int) bool {
	return el[i].name < el[j].name
}

// Swap swaps two elements
func (el Elements) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}

// Sort sorts the element list
func (el Elements) Sort() {
	sort.Sort(el)
}
