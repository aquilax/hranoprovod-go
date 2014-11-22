package main

import (
	"github.com/Hranoprovod/parser"
)

// NodeList contains list of general nodes
type NodeList map[string]*parser.Node

// NewNodeList creates new list of general nodes
func NewNodeList() *NodeList {
	return &NodeList{}
}

func (db *NodeList) push(node *parser.Node) {
	(*db)[(*node).Header] = node
}
