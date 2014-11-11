package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResolver(t *testing.T) {
	Convey("Given nodes database and reslover", t, func() {
		nl := NewNodeList()
		node1 := NewNode("node1")
		node1.elements.add("element1", 100)
		node1.elements.add("element2", 200)
		nl.push(node1)
		node2 := NewNode("node2")
		node2.elements.add("node1", 2)
		nl.push(node2)
		resolver := NewResolver(nl, 1)
		Convey("Resolve resolves the database", func() {
			resolver.resolve()
			Convey("Elements are resolved", func() {
				n1 := (*nl)["node1"]
				So((*n1.elements)[0].name, ShouldEqual, "element1")
				So((*n1.elements)[0].val, ShouldEqual, 100)
				So((*n1.elements)[1].name, ShouldEqual, "element2")
				So((*n1.elements)[1].val, ShouldEqual, 200)
				n2 := (*nl)["node2"]
				So((*n2.elements)[0].name, ShouldEqual, "element1")
				So((*n2.elements)[0].val, ShouldEqual, 200)
				So((*n2.elements)[1].name, ShouldEqual, "element2")
				So((*n2.elements)[1].val, ShouldEqual, 400)
			})
		})
	})
}
