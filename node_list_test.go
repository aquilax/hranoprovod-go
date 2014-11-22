package main

import (
	"github.com/Hranoprovod/parser"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNodeList(t *testing.T) {
	Convey("Given NodeList", t, func() {
		nl := NewNodeList()
		Convey("Creates new NodeList", func() {
			So(nl != nil, ShouldBeTrue)
		})
		Convey("Adding new node", func() {
			node := parser.NewNode("test")
			nl.push(node)
			Convey("Increases the number of nodes in the list", func() {
				So(len(*nl), ShouldEqual, 1)
			})
		})
	})
}
