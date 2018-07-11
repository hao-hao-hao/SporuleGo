package common

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCheckNil(t *testing.T) {
	convey.Convey("Testing to see if CheckNil can really check nil", t, func() {
		convey.Convey("Testing passing two variables , it should return true", func() {
			a := "a"
			b := "b"
			convey.So(CheckNil(a, b), convey.ShouldBeTrue)
		})
		convey.Convey("Testing pass string variable and nil, it should return false", func() {
			a := "a"
			convey.So(CheckNil(a, nil), convey.ShouldBeFalse)
		})
		convey.Convey("Testing an empty instance of struct, it should return true", func() {
			a := &testStruct{}
			convey.So(CheckNil(a), convey.ShouldBeTrue)
		})
		convey.Convey("Empty slice should return false", func() {
			a := []string{}
			b := "b"
			convey.So(CheckNil(a, b), convey.ShouldBeFalse)
		})
		convey.Convey("Not empty slice should return true", func() {
			a := []string{"test"}
			b := "b"
			convey.So(CheckNil(a, b), convey.ShouldBeTrue)
		})
		convey.Convey("Not empty map should return true", func() {
			a := make(map[string]int)
			a["hello"] = 3
			b := "b"
			convey.So(CheckNil(a, b), convey.ShouldBeTrue)
		})
		convey.Convey("Empty map should return false", func() {
			a := make(map[string]int)
			b := "b"
			convey.So(CheckNil(a, b), convey.ShouldBeFalse)
		})
	})
}

func TestGenerateRandomString(t *testing.T) {
	convey.Convey("Random number should genenrate 2 different numebrs", t, func() {
		a := GenerateRandomString()
		b := GenerateRandomString()
		convey.So(a, convey.ShouldNotEqual, b)
	})
}

func TestGetError(t *testing.T) {
	convey.Convey("Testing Get Error function", t, func() {
		convey.Convey("Empty error or nil should return empty string", func() {
			err := errors.New("")
			convey.So(GetError(err), convey.ShouldEqual, "")
			convey.So(GetError(nil), convey.ShouldEqual, "")
		})
		convey.Convey("Error should return error message", func() {
			err := errors.New("error message")
			convey.So(GetError(err), convey.ShouldEqual, "error message")
		})
	})
}

func TestStructToBSON(t *testing.T) {
	convey.Convey("Testing StructToBSON", t, func() {
		convey.Convey("Struct with sub struct arrays should return ")
	})
}
