package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewLogNode(t *testing.T) {
	Convey("Given NewLogNode", t, func() {
		now := time.Now()
		elements := NewElements()
		elements.add("test", 1.22)
		logNode := NewLogNode(now, elements)
		Convey("Creates new log node with the proper fields", func() {
			So(logNode.time.Equal(now), ShouldBeTrue)
			So(logNode.elements, ShouldEqual, elements)
			So((*logNode.elements)[0].name, ShouldEqual, "test")
			So((*logNode.elements)[0].val, ShouldEqual, 1.22)
		})
	})
}
