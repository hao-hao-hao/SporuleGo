package common

import (
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestHTTPResponse(t *testing.T) {
	convey.Convey("Testing HTTPResponse", t, func() {
		c := &gin.Context{}
		monkey.PatchInstanceMethod(reflect.TypeOf(c), "JSON", func(_ *gin.Context, _ int, body interface{}) {
			c.Set("Test", body)
		})
		convey.Convey("Contains error, It should contain error in the result without results", func() {
			monkey.Patch(CheckNil, func(_ ...interface{}) bool {
				return true
			})
			results := &gin.H{"abc": "abc"}
			err := "hello"
			HTTPResponse(c, 200, results, err)
			output, _ := c.Get("Test")
			convey.So(output, convey.ShouldContainKey, "errors")
			convey.So(output, convey.ShouldNotContainKey, "data")
		})

		convey.Convey("Does not contain error, It should contain data without error", func() {
			monkey.Patch(CheckNil, func(_ ...interface{}) bool {
				return false
			})
			results := &gin.H{"abc": "abc"}
			err := "hello"
			HTTPResponse(c, 200, results, err)
			output, _ := c.Get("Test")
			convey.So(output, convey.ShouldContainKey, "data")
			convey.So(output, convey.ShouldNotContainKey, "errors")
		})
	})
}
