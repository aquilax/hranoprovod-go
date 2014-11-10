package main

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