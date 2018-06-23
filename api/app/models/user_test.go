package models

import (
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewUser(t *testing.T) {
	convey.Convey("Testing NewUser", t, func() {
		convey.Convey("Has Nil Values: Should return error without the result", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return false
			})
			result, err := NewUser("", "", "", []Role{})
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Does not have Nil Values: Should return result without the error", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return true
			})
			result, err := NewUser("", "", "", []Role{})
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}
