package common

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/gin-gonic/gin"
)

func TestHTTPResponse(t *testing.T) {
	convey.Convey("Testing Http Response", t, func() {
		c := &gin.Context{}
		convey.Convey("Check if the error is ")
		c.JSON = func(abc int, obj interface{}) {}
	})

}
