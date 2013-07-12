package main

import "fmt"

type Node struct {
	name     string
	elements Elements
}

type NodeList map[string]*Node

func (node *Node) Print() {
	fmt.Printf("name: %s\n", node.name)
}

//Used for debugging
func (nl *NodeList) Print() {
	for _, node := range *(nl) {
		fmt.Println(node.name)
		for _, e := range node.elements {
			fmt.Printf("\t%s : %0.2f\n", e.name, e.val)
		}
	}
}
