package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBreakingError(t *testing.T) {
	Convey("Given new breaking error", t, func() {
		err := NewBreakingError("test", 1)
		Convey("Error is of the right type", func() {
			So(err.Error(), ShouldEqual, "test")
		})
	})
}
