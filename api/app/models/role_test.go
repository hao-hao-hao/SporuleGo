package models

import (
	"errors"
	"sporule/api/app/modules/test"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
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

func TestInsert(t *testing.T) {
	//apply patches
	helper := &test.Helper{}
	helper.PatchResouces()
	//custom patches
	helper.AddPatches(
		monkey.Patch(GetRoleByName, func(name string) (*Role, error) {
			//simulate the database only contain Admin in Role
			if name == "Admin" {
				return &Role{Name: "Admin"}, nil
			}
			return &Role{}, errors.New("Couldn't find the role")
		}))
	defer helper.Unpatch()
	convey.Convey("Testing Role.Insert", t, func() {
		convey.Convey("Inserting the role that is already in DB, the error message should not be nil", func() {
			role := &Role{Name: "Admin"}
			err := role.Insert()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Inserting the role that is not in DB, the error message should  be nil", func() {
			role := &Role{Name: "Memeber"}
			err := role.Insert()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
