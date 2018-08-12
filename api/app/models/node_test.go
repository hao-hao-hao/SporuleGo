package models

import (
	"sporule/api/app/modules/test"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestNewNode(t *testing.T) {
	convey.Convey("Testing NewNode", t, func() {
		convey.Convey("Has Nil Values: Should return error without the result", func() {
			result, err := NewNode("", NodeTemplate{}, Node{}, []Role{Role{Name: "ABC"}})
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Does not have Nil Values: Should return result without the error. The result should contain TemplateID and PermittedRoleID", func() {
			result, err := NewNode("FirstNode", NodeTemplate{ID: "32132"}, Node{ID: "123123"}, []Role{Role{Name: "ABC"}})
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
			convey.So(result.ID, convey.ShouldNotBeNil)
			convey.So(result.TemplateID, convey.ShouldNotBeNil)
			convey.So(result.PermittedRoleIds, convey.ShouldNotBeNil)
		})
	})
}

func TestNodeInsert(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	defer helper.Unpatch()
	convey.Convey("Testing Node.Insert", t, func() {
		convey.Convey("Has Nil values should return an error", func() {
			node := &Node{ID: "123", Name: ""}
			err := node.Insert("123")
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("It contains all necessary information should return no error, and createddate and modifieddate should equal to now", func() {
			node := &Node{ID: "123", Name: "Test", TemplateID: "123", PermittedRoleIds: []bson.ObjectId{"123", "223"}}
			err := node.Insert("123")
			convey.So(err, convey.ShouldBeNil)
			convey.So(node.CreatedDate, convey.ShouldEqual, time.Now())
			convey.So(node.ModifiedDate, convey.ShouldEqual, time.Now())
		})
	})
}

func TestNodeUpdate(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	defer helper.Unpatch()

	convey.Convey("Testing Node.Update", t, func() {
		convey.Convey("Has Nil values should return an error", func() {
			node := &Node{ID: "123", Name: ""}
			err := node.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("It contains all necessary information should return no error, and it should modifieddate should equal to now", func() {
			node := &Node{ID: "123", Name: "Test", TemplateID: "123", PermittedRoleIds: []bson.ObjectId{"123", "223"}, CreatorID: "123123"}
			err := node.Update()
			convey.So(err, convey.ShouldBeNil)
			convey.So(node.ModifiedDate, convey.ShouldEqual, time.Now())
		})
	})
}
