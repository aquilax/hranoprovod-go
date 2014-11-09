package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMytrim (t *testing.T) {
	Convey("Given whitespaced strings", t, func() {
		testCases := []string{
			"trim",
			" trim",
			"	trim",
			" trim	",
		}
		Convey("Clears the white space", func() {
			for _, testCase := range testCases {
				So(mytrim(testCase), ShouldEqual, "trim")
			}
		})
	})
}

func TestParseTime (t *testing.T) {
	Convey("Givent correct time string", t, func() {
		stringTime := "1980/12/17"
		Convey("Returns correct time", func() {
			result, err := parseTime(stringTime)
			So(result.Equal(time.Date(1980, 12, 17, 0, 0 ,0 , 0, time.UTC)), ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
	Convey("Givent incorrect time string", t, func() {
		stringTime := "19/12/17"
		Convey("Returns error", func() {
			_, err := parseTime(stringTime)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestIsGoodDate(t *testing.T) {
	Convey("For given date", t, func() {
		d := time.Date(1980, 12, 17, 0, 0 ,0 , 0, time.UTC)
		Convey("If beginning is equal or before the date", func() {
			testCases := []time.Time{
				time.Date(1980, 12, 17, 0, 0 ,0 , 0, time.UTC),
				time.Date(1980, 12, 16, 0, 0 ,0 , 0, time.UTC),
			}
			Convey("Should return true", func(){
				for _, testCase := range testCases {
					So(isGoodDate(d, testCase, dateBeginning), ShouldBeTrue)
				}
				
			})
		})
		Convey("If beginning is after the date", func() {
			b := time.Date(1980, 12, 18, 0, 0 ,0 , 0, time.UTC)
			Convey("Should retutn false", func() {
				So(isGoodDate(d, b, dateBeginning), ShouldBeFalse)
			})
		})
		Convey("If end is equal or after the date", func() {
			testCases := []time.Time{
				time.Date(1980, 12, 17, 0, 0 ,0 , 0, time.UTC),
				time.Date(1980, 12, 18, 0, 0 ,0 , 0, time.UTC),
			}
			Convey("Should return true", func(){
				for _, testCase := range testCases {
					So(isGoodDate(d, testCase, dateEnd), ShouldBeTrue)
				}
				
			})
		})
		Convey("If end is before the date", func() {
			b := time.Date(1980, 12, 16, 0, 0 ,0 , 0, time.UTC)
			Convey("Should retutn false", func() {
				So(isGoodDate(d, b, dateEnd), ShouldBeFalse)
			})
		})

	})
}