package main

// Node contains general node data
type Node struct {
	header   string
	elements *Elements
}

// NewNode creates new geneal node
func NewNode(header string) *Node {
	return &Node{
		header,
		NewElements(),
	}
}

// NodeList contains list of general nodes
type NodeList map[string]*Node

// NewNodeList creates new list of general nodes
func NewNodeList() *NodeList {
	return &NodeList{}
}

func (db *NodeList) push(node *Node) {
	(*db)[(*node).header] = node
}
