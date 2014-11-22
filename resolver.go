package main

import (
	"github.com/Hranoprovod/parser"
)

// Resolver contains the resolver data
type Resolver struct {
	db       *NodeList
	maxDepth int
}

// NewResolver creates new resolver
func NewResolver(db *NodeList, maxDepth int) *Resolver {
	return &Resolver{db, maxDepth}
}

func (r *Resolver) resolve() {
	for name := range *r.db {
		r.resolveNode(name, 0)
	}
}

func (r *Resolver) resolveNode(name string, level int) {
	if level >= r.maxDepth {
		return
	}

	node, exists := (*r.db)[name]
	if !exists {
		return
	}

	nel := parser.NewElements()

	for _, e := range *node.Elements {
		r.resolveNode(e.Name, level+1)
		snode, exists := (*r.db)[e.Name]
		if exists {
			nel.SumMerge(snode.Elements, e.Val)
		} else {
			var tm parser.Elements
			tm.Add(e.Name, e.Val)
			nel.SumMerge(&tm, 1)
		}
	}
	nel.Sort()
	(*r.db)[name].Elements = nel
}
