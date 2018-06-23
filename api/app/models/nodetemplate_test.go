package models

import (
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewNodeTemplate(t *testing.T) {
	convey.Convey("Testing NewNodeTemplate", t, func() {
		convey.Convey("Has Nil Values: Should return error without the result", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return false
			})
			result, err := NewNodeTemplate("", "", []Field{})
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Does not have Nil Values: Should return result without the error", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return true
			})
			result, err := NewNodeTemplate("", "", []Field{})
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}
