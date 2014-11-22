package main

import (
	"github.com/Hranoprovod/parser"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewLogNode(t *testing.T) {
	Convey("Given NewLogNode", t, func() {
		now := time.Now()
		elements := parser.NewElements()
		elements.Add("test", 1.22)
		logNode := NewLogNode(now, elements)
		Convey("Creates new log node with the proper fields", func() {
			So(logNode.time.Equal(now), ShouldBeTrue)
			So(logNode.elements, ShouldEqual, elements)
			So((*logNode.elements)[0].Name, ShouldEqual, "test")
			So((*logNode.elements)[0].Val, ShouldEqual, 1.22)
		})
	})
}
