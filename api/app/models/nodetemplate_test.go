package models

import (
	"sporule/api/app/modules/test"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewNodeTemplate(t *testing.T) {
	convey.Convey("Testing NewNodeTemplate", t, func() {
		convey.Convey("Name is nil: Should return error without the result", func() {
			result, err := NewNodeTemplate("", nil)
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Name is not nil: Should return result without the error", func() {
			result, err := NewNodeTemplate("Shop", nil)
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}

func TestNodeTemplateInsert(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	defer helper.Unpatch()
	convey.Convey("Testing NodeTemplate.Insert", t, func() {
		convey.Convey("Nil ID will result an error", func() {
			nodeTemplate := &NodeTemplate{Name: "TemplateA"}
			err := nodeTemplate.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Node Template with ID and name will not return an error", func() {
			nodeTemplate := &NodeTemplate{ID: "123", Name: "TemplateA"}
			err := nodeTemplate.Insert()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
