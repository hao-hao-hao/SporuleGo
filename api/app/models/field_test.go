package models

import (
	"errors"
	"sporule/api/app/modules/test"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestNewField(t *testing.T) {
	convey.Convey("Testing NewField constructor", t, func() {
		convey.Convey("Name is empty, so result should be nil, error message should not be nil", func() {
			result, err := NewField("", "TextBox")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(result, convey.ShouldBeNil)
		})
		convey.Convey("Both Name and Field is not empty, so result should not be nil and error should be nil", func() {
			result, err := NewField("Name", "TextBox")
			convey.So(result, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFieldInsert(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	helper.AddPatches(
		monkey.Patch(GetFieldByName, func(name string) (*Field, error) {
			//simulate the database only contain Admin in Role
			if name == "FieldA" {
				return &Field{Name: "FieldA"}, nil
			}
			return &Field{}, errors.New("Couldn't find the role")
		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing Insert", t, func() {
		convey.Convey("Field with empty name should return an error", func() {
			field := &Field{Type: "TextBox"}
			err := field.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Field name is already exist in database should return an error", func() {
			field := &Field{Name: "FieldA"}
			err := field.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Field Name is not exist should return no error", func() {
			field := &Field{Name: "FieldB", Type: "ABC", ID: "123"}
			err := field.Insert()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFieldUpdate(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	helper.AddPatches(
		monkey.Patch(GetFieldByName, func(name string) (*Field, error) {
			//simulate the database only contain Admin in Role
			if name == "FieldA" || name == "FieldB" {
				return &Field{Name: "FieldA"}, nil
			}
			return &Field{}, errors.New("Couldn't find the role")
		}),
		monkey.Patch(GetFieldByID, func(id bson.ObjectId) (*Field, error) {
			return &Field{Name: "FieldB"}, errors.New("Couldn't find the role")
		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing Insert", t, func() {
		convey.Convey("Field with empty name should return an error", func() {
			field := &Field{Type: "TextBox"}
			err := field.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Field name is already exist in database should return an error", func() {
			field := &Field{Name: "FieldA"}
			err := field.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Field Name is not exist should return no error", func() {
			field := &Field{Name: "FieldB", Type: "ABC", ID: "123"}
			err := field.Update()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
