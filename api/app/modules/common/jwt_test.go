package common

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetIDInHeader(t *testing.T) {
	convey.Convey("Testing GetIDInHeader", t, func() {
		Enums.loadOtherEnums()
		convey.Convey("It should return the id abc", func() {
			c := &gin.Context{}
			c.Set(Enums.Others.IDInHeader, "abc")
			monkey.Patch(CheckNil, func(_ ...interface{}) bool {
				return true
			})
			convey.So(GetIDInHeader(c), convey.ShouldEqual, "abc")
		})

		convey.Convey("It should throw abort error if id header is not set", func() {
			c := &gin.Context{}
			monkey.Patch(CheckNil, func(_ ...interface{}) bool {
				return false
			})
			monkey.Patch(HTTPResponse401, func(_ *gin.Context) {
			})
			convey.So(GetIDInHeader(c), convey.ShouldEqual, "")
		})

	})
}
