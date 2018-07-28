package apicontrollers

import (
	"errors"
	"reflect"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestAddUser(t *testing.T) {
	convey.Convey("Testing AddUser", t, func() {
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.Patch(models.NewUser, func(_, _, _ string, _ []models.Role) (*models.User, error) {
			return &models.User{}, nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "Register", func(_ *models.User) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(c *gin.Context, _ interface{}) error {
			return nil
		})

		convey.Convey("Binding JSON Failure should return error message", func() {
			c := &gin.Context{}
			monkey.PatchInstanceMethod(reflect.TypeOf(c), "BindJSON", func(c *gin.Context, _ interface{}) error {
				return errors.New("error")
			})
			AddUser(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Lack of information should return error message", func() {
			c := &gin.Context{}

			monkey.Patch(models.NewUser, func(_, _, _ string, _ []models.Role) (*models.User, error) {
				return nil, errors.New("errors")
			})
			AddUser(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("User exist should return error message", func() {
			c := &gin.Context{}

			monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "Register", func(_ *models.User) error {
				return errors.New("user exist")
			})
			AddUser(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Perfect information without duplicate users, it should return the user obejct", func() {
			c := &gin.Context{}

			monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "Register", func(_ *models.User) error {
				return nil
			})
			AddUser(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})

	})
}

func TestAddField(t *testing.T) {
	convey.Convey("Testing AddField", t, func() {
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.Patch(models.NewField, func(_, _ string) (*models.Field, error) {
			return &models.Field{}, nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.Field{}), "Insert", func(_ *models.Field) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(c *gin.Context, _ interface{}) error {
			return nil
		})

		convey.Convey("Binding JSON Failure should return error message", func() {
			c := &gin.Context{}
			monkey.PatchInstanceMethod(reflect.TypeOf(c), "BindJSON", func(c *gin.Context, _ interface{}) error {
				return errors.New("error")
			})
			AddField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Lack of information should return error message", func() {
			c := &gin.Context{}

			monkey.Patch(models.NewField, func(_, _ string) (*models.Field, error) {
				return nil, errors.New("errors")
			})
			AddField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Field exist should return error message", func() {
			c := &gin.Context{}

			monkey.PatchInstanceMethod(reflect.TypeOf(&models.Field{}), "Insert", func(_ *models.Field) error {
				return errors.New("Field exist")
			})
			AddField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Perfect information without duplicate fields, it should return the field obejct", func() {
			c := &gin.Context{}
			monkey.PatchInstanceMethod(reflect.TypeOf(&models.Field{}), "Insert", func(_ *models.Field) error {
				return nil
			})
			AddField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})

	})
}

func TestDeleteField(t *testing.T) {
	convey.Convey("Testing DeleteField", t, func() {
		monkey.Patch(common.HTTPResponse404, func(c *gin.Context) {
			c.Set("results", "error")
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, _ *gin.H, _ string) {
			c.Set("results", "result")
		})
		monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
			return "id12312321", nil
		})
		monkey.Patch(models.DeleteField, func(_ bson.ObjectId) error {
			return nil
		})
		convey.Convey("Invalid ID should raise an error", func() {
			patch := monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
				return "", errors.New("errors")
			})
			c := &gin.Context{}
			DeleteField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Non exist field will return error", func() {
			patch := monkey.Patch(models.DeleteField, func(_ bson.ObjectId) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			DeleteField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Correct Existing Field ID should return without error", func() {
			c := &gin.Context{}
			DeleteField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}

func TestUpdateField(t *testing.T) {
	convey.Convey("Testing UpdateField", t, func() {
		monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
			return "id123123", nil
		})
		monkey.Patch(common.HTTPResponse404, func(c *gin.Context) {
			c.Set("results", "error")
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, _ *gin.H, _ string) {
			c.Set("results", "result")
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.Field{}), "Update", func(_ *models.Field, _ bson.ObjectId) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(_ *gin.Context, _ interface{}) error {
			return nil
		})
		convey.Convey("Invalid ID would return error", func() {
			patch := monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
				return "", errors.New("error")
			})
			c := &gin.Context{}
			UpdateField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Unfound Field would return error", func() {
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(_ *gin.Context, _ interface{}) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			UpdateField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Correct Existing Field ID would return result", func() {
			c := &gin.Context{}
			UpdateField(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}
