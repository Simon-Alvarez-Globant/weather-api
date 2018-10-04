package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetconfig(t *testing.T) {

	Convey("The GetConfigs func", t, func() {
		r := GetConfigs()
		var configType Config
		Convey("should return a Config struct type", func() {
			So(r, ShouldHaveSameTypeAs, configType)
		})
	})
}
