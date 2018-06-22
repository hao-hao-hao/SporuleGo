package common

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLoadConfiguration(t *testing.T) {
	convey.Convey("Testing LoadConfiguration", t, func() {
		convey.Convey("wrong file path should return an error", func() {
			convey.So(Config.LoadConfiguration("../config/test.json"), convey.ShouldNotBeNil)
		})
		convey.Convey("Correct file should not return error", func() {
			convey.So(Config.LoadConfiguration("../../../config/test.json"), convey.ShouldBeNil)
		})
	})
}
