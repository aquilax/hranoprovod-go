package main

import(
  "fmt"
  "sort"
)

type Options struct{
  unresolved bool
  version bool
  help bool
  totals bool
  csv bool
  single_element string
  single_food string
  beginning string
  end string

  log_file_name string
  database_file_name string
}

type Element struct {
  name string
  val float32
}

type Elements []Element

type Node struct {
  name string
  elements Elements
}

type NodeList map[string] *Node

type Accumulator map[string] [2]float32

//Used for debugging
func (nl * NodeList) Print () {
  for _, node := range *(nl) {
    fmt.Println(node.name);
    for _, e := range node.elements{
      fmt.Printf("\t%s : %0.2f\n", e.name, e.val)
    }
  }
}

func (el *Elements) Index (name string) (int, bool) {
  for n, e := range *el {
    if e.name == name {
      return n, true
    }
  }
  return 0, false
}

func (el *Elements) Add (name string, val float32) {
  var e Element
  e.name = name
  e.val = val
  *el = append(*el, e);
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
  sort.Sort(el);
}



