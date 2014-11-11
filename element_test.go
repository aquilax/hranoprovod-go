package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestElement(t *testing.T) {
	Convey("NewElement", t, func() {
		el := NewElement("test", 10)
		Convey("Creates new element", func() {
			So(el.name, ShouldEqual, "test")
			So(el.val, ShouldEqual, 10)
		})
	})
}

func TestElements(t *testing.T) {
	Convey("Given Elements", t, func() {
		el := NewElements()
		Convey("Calling Add", func() {
			el.add("test", 10)
			Convey("Adds the element to the list", func() {
				So(el.Len(), ShouldEqual, 1)
			})
			Convey("After adding more elements", func() {
				el.add("test3", 13)
				el.add("test2", 12)
				el.add("test1", 11)
				Convey("Calling Index on present element", func() {
					index, found := el.index("test2")
					Convey("Returns the correct index", func() {
						So(index, ShouldEqual, 2)
					})
					Convey("Returns positive found", func() {
						So(found, ShouldBeTrue)
					})
				})
				Convey("Calling Index on missing element", func() {
					_, found := el.index("test111")
					Convey("Returns not found", func() {
						So(found, ShouldBeFalse)
					})
				})
				Convey("After Sort", func() {
					el.Sort()
					Convey("Elements are sorted", func() {
						index, _ := el.index("test3")
						So(index, ShouldEqual, 3)
						index2, _ := el.index("test1")
						So(index2, ShouldEqual, 1)
					})
				})
				Convey("Having second set of elements", func() {
					el2 := NewElements()
					el2.add("test3", 113)
					el2.add("test2", 112)
					el2.add("test1", 111)
					el2.add("test4", 444)
					Convey("SumMerge with coef 2", func() {
						el.sumMerge(el2, 2)
						Convey("Returns correct elements", func() {
							index, found := el.index("test1")
							So(found, ShouldBeTrue)
							So(index, ShouldEqual, 3)
							So((*el)[index].val, ShouldEqual, 233)
						})
						Convey("New elements are added", func() {
							index, found := el.index("test4")
							So(found, ShouldBeTrue)
							So((*el)[index].val, ShouldEqual, 888)
						})
					})
				})
			})
		})
	})
}
