package main

const MAX_LEVEL = 9

type Resolver struct {
	db *NodeList
}

func NewResolver(db *NodeList) *Resolver {
	return &Resolver{db}
}

func (r *Resolver) resolve() {
	for name, _ := range *r.db {
		r.resolveNode(name, 0)
	}
}

func (r *Resolver) resolveNode(name string, level int) {
	if level > MAX_LEVEL {
		return
	}

	node, exists := (*r.db)[name]
	if !exists {
		return
	}

	var nel Elements

	for _, e := range *node.elements {
		r.resolveNode(e.name, level+1)
		snode, exists := (*r.db)[e.name]
		if exists {
			nel.SumMerge(snode.elements, e.val)
		} else {
			var tm Elements
			tm.Add(e.name, e.val)
			nel.SumMerge(&tm, 1)
		}
	}
	nel.Sort()
	*(*r.db)[name].elements = nel
}
