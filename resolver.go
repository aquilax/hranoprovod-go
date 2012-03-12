package main

func (els *Elements) SumMerge(left *Elements, coef float32) {
	for _, v := range *left {
		ndx, exists := (*els).Index(v.name)
		if exists {
			(*els)[ndx].val += v.val * coef
		} else {
			(*els).Add(v.name, v.val*coef)
		}
	}
}

func (db *NodeList) ResolveNode(name string, level int) {
	if level > 9 {
		return
	}

	node, exists := (*db)[name]
	if !exists {
		return
	}

	var nel Elements

	for _, e := range node.elements {
		db.ResolveNode(e.name, level+1)
		snode, exists := (*db)[e.name]
		if exists {
			nel.SumMerge(&snode.elements, e.val)
		} else {
			var tm Elements
			tm.Add(e.name, e.val)
			nel.SumMerge(&tm, 1)
		}
	}
	nel.Sort()
	(*db)[name].elements = nel
}

func (db *NodeList) Resolve() {
	for name, _ := range *db {
		db.ResolveNode(name, 0)
	}
}
