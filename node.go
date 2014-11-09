package main

import "fmt"

type Node struct {
	name     string
	elements Elements
}

func NewNode() *Node {
	return &Node{}
}

type NodeList map[string]*Node

func NewNodeList() *NodeList {
	return &NodeList{}
}

func (db *NodeList) Push(node *Node) {
	(*db)[(*node).name] = node
}

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
