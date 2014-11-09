package main

import "fmt"

type Node struct {
	header   string
	elements *Elements
}

func NewNode() *Node {
	return &Node{
		"",
		NewElements(),
	}
}

type NodeList map[string]*Node

func NewNodeList() *NodeList {
	return &NodeList{}
}

func (db *NodeList) Push(node *Node) {
	(*db)[(*node).header] = node
}

func (node *Node) Print() {
	fmt.Printf("header: %s\n", node.header)
}

//Used for debugging
func (nl *NodeList) Print() {
	for _, node := range *(nl) {
		fmt.Println(node.header)
		for _, e := range *node.elements {
			fmt.Printf("\t%s : %0.2f\n", e.name, e.val)
		}
	}
}
