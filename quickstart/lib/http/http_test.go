package http

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLib(t *testing.T) {
	Convey("Test windDirection func \n", t, func() {
		r := windDirection(360)
		Convey("Should return a string", func() {
			So(r, ShouldHaveSameTypeAs, "")
		})
		Convey("Should return north by default", func() {
			So(r, ShouldEqual, "north")
		})
		Convey("Should return a string for each angle range", func() {
			So(windDirection(0), ShouldEqual, "north")
			So(windDirection(22.5), ShouldEqual, "north-northeast")
			So(windDirection(45), ShouldEqual, "northeast")
			So(windDirection(67.5), ShouldEqual, "east-northeast")
			So(windDirection(90), ShouldEqual, "east")
			So(windDirection(112.5), ShouldEqual, "east-southeast")
			So(windDirection(135), ShouldEqual, "southeast")
			So(windDirection(157.5), ShouldEqual, "south-southeast")
			So(windDirection(180), ShouldEqual, "south")
			So(windDirection(202.5), ShouldEqual, "south-southwest")
			So(windDirection(225), ShouldEqual, "southwest")
			So(windDirection(247.5), ShouldEqual, "west-southwest")
			So(windDirection(270), ShouldEqual, "west")
			So(windDirection(292.5), ShouldEqual, "west-northwest")
			So(windDirection(315), ShouldEqual, "northwest")
			So(windDirection(337.5), ShouldEqual, "north-northweast")
		})
	})

	Convey("Test windCondition function \n", t, func() {
		r := windCondition(0)
		Convey("Should return a string", func() {
			So(r, ShouldHaveSameTypeAs, "")
		})
		Convey("Should return 'Be careful, crazy wind outside!' by default", func() {
			So(windCondition(1000), ShouldEqual, "Be careful, crazy wind outside!")
		})
		Convey("Should return a different string for each possible case", func() {
			So(windCondition(0), ShouldEqual, "Gentle breeze")
			So(windCondition(5), ShouldEqual, "Windy")
		})
	})

}
