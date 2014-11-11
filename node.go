package main

type Node struct {
	header   string
	elements *Elements
}

func NewNode(header string) *Node {
	return &Node{
		header,
		NewElements(),
	}
}

type NodeList map[string]*Node

func NewNodeList() *NodeList {
	return &NodeList{}
}

func (db *NodeList) push(node *Node) {
	(*db)[(*node).header] = node
}
