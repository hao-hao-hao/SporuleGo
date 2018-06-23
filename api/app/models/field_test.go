package models

import (
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewField(t *testing.T) {
	convey.Convey("Testing NewField constructor", t, func() {
		convey.Convey("Name is empty, so result should be nil, error message should not be nil", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return false
			})
			result, err := NewField("", "Other")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(result, convey.ShouldBeNil)
		})
		convey.Convey("Both Name and Field is not empty, so result should not be nil and error should be nil", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return true
			})
			result, err := NewField("", "Other")
			convey.So(result, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
