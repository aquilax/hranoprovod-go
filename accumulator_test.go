package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAccumulator(t *testing.T) {
	Convey("Given the acccumulator", t, func() {
		acc := NewAccumulator()
		Convey("When a positive element is added", func() {
			acc.add("test", 1.22)
			Convey("It should go to the positive acccumulator", func() {
				So((*acc)["test"][accPos], ShouldEqual, 1.22)
			})
			Convey("When a positive value is added to the same key", func() {
				acc.add("test", 2.33)
				Convey("It is accumulated in the positive register", func() {
					So((*acc)["test"][accPos], ShouldEqual, 3.55)
				})
			})
			Convey("When a negative value is added to the same key", func() {
				acc.add("test", -2.33)
				Convey("It is accumulated in the positive register", func() {
					So((*acc)["test"][accNeg], ShouldEqual, -2.33)
					So((*acc)["test"][accPos], ShouldEqual, 1.22)
				})
			})
			Convey("When negative element is added", func() {
				acc.add("test2", -1.32)
				Convey("It should go to the negative acccumulator", func() {
					So((*acc)["test2"][accNeg], ShouldEqual, -1.32)
				})
			})
		})
	})
}
