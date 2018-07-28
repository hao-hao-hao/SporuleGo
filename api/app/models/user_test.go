package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"sporule/api/app/modules/test"
	"testing"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"

	"gopkg.in/mgo.v2/bson"
)

func TestNewUser(t *testing.T) {
	//patches
	helper := &test.Helper{}
	helper.PatchResouces()
	defer helper.Unpatch()
	convey.Convey("Testing NewUser", t, func() {
		convey.Convey("email/password is Nil : Should return error without the result", func() {
			user, err := NewUser("", "")
			convey.So(user, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("email/password is not nil: Should return result without the error. The user object should contain ID,email,password,failedlogin,isdisabled and tokensalt", func() {
			user, err := NewUser("abc@gmail.com", "1q2w3e4r")
			convey.So(err, convey.ShouldBeNil)
			convey.So(common.CheckNil(user.ID, user.Email, user.Password, user.FailedLogin, user.IsDisabled, user.TokenSalt), convey.ShouldBeTrue)
		})
	})
}

func TestRegister(t *testing.T) {
	//patches
	helper := &test.Helper{}
	helper.PatchResouces()
	helper.AddPatches(
		monkey.Patch(GetUserByEmail, func(email string) (*User, error) {
			if email == "abc@gmail.com" {
				password, _ := common.EncryptPassword("1q2w3e4r")
				return &User{Email: email, Password: password}, nil
			}
			return &User{}, errors.New(common.Enums.ErrorMessages.PageNotFound)
		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing Register", t, func() {
		convey.Convey("email/password/id is nil: should return error", func() {
			user := &User{}
			err := user.Register()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("email is already exist: should not reutrn error", func() {
			user := &User{ID: "123", Email: "abc@gmail.com", Password: "1q2w3e4r"}
			err := user.Register()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("email/password/id is not nil: should not reutrn error", func() {
			user := &User{ID: "123", Email: "bbc@gmail.com", Password: "1q2w3e4r"}
			err := user.Register()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestUserUpdate(t *testing.T) {
	//patches
	helper := &test.Helper{}
	helper.PatchResouces()
	helper.AddPatches(
		//These patches is to simulate the condition of if user.IsExist() && user.Email != tempUser.Email
		monkey.Patch(GetUserByEmail, func(email string) (*User, error) {
			if email == "abc@gmail.com" || email == "bbc@gmail.com" {
				password, _ := common.EncryptPassword("1q2w3e4r")
				return &User{Email: email, Password: password}, nil
			}
			return &User{}, errors.New(common.Enums.ErrorMessages.PageNotFound)
		}),
		monkey.Patch(GetUserByID, func(id bson.ObjectId) (*User, error) {
			return &User{Email: "bbc@gmail.com"}, nil

		}),
	)
	defer helper.Unpatch()
	convey.Convey("Testing User Update", t, func() {
		convey.Convey("email/password/id is nil: should return error", func() {
			user := &User{}
			err := user.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("email is already exist: should not reutrn error", func() {
			user := &User{ID: "123", Email: "abc@gmail.com", Password: "1q2w3e4r"}
			err := user.Update()
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("updating user without email change, it should not reutrn error", func() {
			user := &User{ID: "123", Email: "bbc@gmail.com", Password: "1q2w3e4r"}
			err := user.Update()
			convey.So(err, convey.ShouldBeNil)
		})
		convey.Convey("updating user with new email, it should not reutrn error", func() {
			user := &User{ID: "123", Email: "bbcd@gmail.com", Password: "1q2w3e4r"}
			err := user.Update()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestVerifyUser(t *testing.T) {
	//patches
	helper := &test.Helper{}
	helper.PatchResouces()
	helper.AddPatches(
		monkey.Patch(GetUserByEmail, func(email string) (*User, error) {
			if email == "abc@gmail.com" {
				password, _ := common.EncryptPassword("1q2w3e4r")
				return &User{Email: email, Password: password}, nil
			}
			return nil, errors.New(common.Enums.ErrorMessages.PageNotFound)
		}),
	)
	//correct username password is abc@gmail.com, 1q2w3e4r
	convey.Convey("Testing VerifyUser", t, func() {
		convey.Convey("wrong email should return nil user and error", func() {
			user, err := VerifyUser("abcd@gmail.com", "1q2w3e4r")
			convey.So(user, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("correct email with wrong password should return nil user and error", func() {
			user, err := VerifyUser("abc@gmail.com", "1q2w3e4r5t")
			convey.So(user, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("correct email and password should return user and nil error", func() {
			user, err := VerifyUser("abc@gmail.com", "1q2w3e4r")
			convey.So(user, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
