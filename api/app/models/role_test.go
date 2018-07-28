package models

import (
	"errors"
	"sporule/api/app/modules/test"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestNewRole(t *testing.T) {
	convey.Convey("Testing NewRole", t, func() {
		convey.Convey("Role Name is Nil: Should return error without the result", func() {
			role, err := NewRole("")
			convey.So(role, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Role Name is Admin: Should return result without the error", func() {
			role, err := NewRole("Admin")
			convey.So(err, convey.ShouldBeNil)
			convey.So(role, convey.ShouldNotBeNil)
		})
	})
}
func TestRoleInsert(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	//custom patches
	helper.AddPatches(
		monkey.Patch(GetRoleByName, func(name string) (*Role, error) {
			//These patches is for testing role.IsExist() && role.Name != tempRole.Name
			//simulate the database only contain Admin in Role
			if name == "Admin" {
				return &Role{Name: "Admin"}, nil
			}
			return &Role{}, errors.New("Couldn't find the role")
		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing Role.Insert", t, func() {
		convey.Convey("Role with empty name should return an error", func() {
			role := &Role{ID: "312"}
			err := role.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Inserting the role that is already in DB, the error message should not be nil", func() {
			role := &Role{Name: "Admin", ID: "123"}
			err := role.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Inserting the role that is not in DB, the error message should  be nil", func() {
			role := &Role{Name: "Memeber", ID: "123"}
			err := role.Insert()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRoleUpdate(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	//custom patches
	helper.AddPatches(
		monkey.Patch(GetRoleByName, func(name string) (*Role, error) {
			//simulate the database only contain Admin in Role
			if name == "Admin" || name == "Member" {
				return &Role{Name: "Admin"}, nil
			}
			return &Role{}, errors.New("Couldn't find the role")
		}),
		monkey.Patch(GetRoleByID, func(id bson.ObjectId) (*Role, error) {
			return &Role{Name: "Member"}, nil
		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing Role.Update", t, func() {
		convey.Convey("Role with empty name should return an error", func() {
			role := &Role{ID: "312"}
			err := role.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Updating the role that is already in DB, the error message should not be nil", func() {
			role := &Role{Name: "Admin", ID: "312"}
			err := role.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Updating the role without name change, the error message should  be nil", func() {
			role := &Role{Name: "Member", ID: "312"}
			err := role.Update()
			convey.So(err, convey.ShouldBeNil)
		})
		convey.Convey("Updating the role that is not in DB, the error message should  be nil", func() {
			role := &Role{Name: "Admin2", ID: "312"}
			err := role.Update()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
