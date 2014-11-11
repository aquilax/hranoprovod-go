package main

type Resolver struct {
	db       *NodeList
	maxDepth int
}

func NewResolver(db *NodeList, maxDepth int) *Resolver {
	return &Resolver{db, maxDepth}
}

func (r *Resolver) resolve() {
	for name, _ := range *r.db {
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

	nel := NewElements()

	for _, e := range *node.elements {
		r.resolveNode(e.name, level+1)
		snode, exists := (*r.db)[e.name]
		if exists {
			nel.sumMerge(snode.elements, e.val)
		} else {
			var tm Elements
			tm.add(e.name, e.val)
			nel.sumMerge(&tm, 1)
		}
	}
	nel.Sort()
	(*r.db)[name].elements = nel
}
