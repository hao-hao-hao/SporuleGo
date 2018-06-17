package common

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCheckNil(t *testing.T) {
	convey.Convey("Testing to see if CheckNil can really check nil", t, func() {

	})
}

func TestGenerateRandomString(t *testing.T) {
	convey.Convey("Random number should genenrate 2 different numebrs", t, func() {
		a := GenerateRandomString()
		b := GenerateRandomString()
		convey.So(a, convey.ShouldNotEqual, b)
	})
}
