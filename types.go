package main

import("fmt")

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

type Elements map[string] float32

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
    for name, val := range node.elements{
      fmt.Printf("\t%s:%f\n", name, val)
    }
  }
}
