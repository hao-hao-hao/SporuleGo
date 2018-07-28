package apicontrollers

import (
	"errors"
	"net/http/httptest"
	"reflect"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestAddRole(t *testing.T) {
	convey.Convey("Testing AddRole", t, func() {
		//Apply Patches
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.Patch(models.NewRole, func(_ string) (*models.Role, error) {
			return &models.Role{}, nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.Role{}), "Insert", func(_ *models.Role) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(c *gin.Context, _ interface{}) error {
			return nil
		})

		//Test scenarios
		convey.Convey("Adding a role without name should return an error", func() {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			//patches to simulate fail binding
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(c), "BindJSON", func(c *gin.Context, _ interface{}) error {
				return errors.New("error")
			})
			defer patch.Unpatch()

			AddRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
		})
		convey.Convey("Lack of information should return error message", func() {
			c := &gin.Context{}
			patch := monkey.Patch(models.NewRole, func(_ string) (*models.Role, error) {
				return nil, errors.New("errors")
			})
			AddRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Field exist should return error message", func() {
			c := &gin.Context{}
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(&models.Role{}), "Insert", func(_ *models.Role) error {
				return errors.New("Role exist")
			})
			AddRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Perfect information without duplicate fields, it should return the field obejct", func() {
			c := &gin.Context{}
			AddRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}

func TestUpdateRole(t *testing.T) {
	convey.Convey("Testing  UpdateRole", t, func() {
		//Apply Patches
		monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
			return "id123", nil
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(c *gin.Context, _ interface{}) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.Role{}), "Update", func(_ *models.Role, _ bson.ObjectId) error {
			return nil
		})

		//Test Scenario
		convey.Convey("Not enough information/id for binding the role object should return error", func() {
			patch := monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
				return "", errors.New("error")
			})
			c := &gin.Context{}
			UpdateRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Role is not exist in the database should return error", func() {
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(&models.Role{}), "Update", func(_ *models.Role, _ bson.ObjectId) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			UpdateRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Correct information has been passed to the db return error", func() {
			c := &gin.Context{}
			UpdateRole(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}

func TestGetRoleByID(t *testing.T) {
	convey.Convey("Testing GetRoleByID", t, func() {
		//Apply Patches
		monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
			return "id123", nil
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.Patch(models.GetRoleByID, func(_ bson.ObjectId) (*models.Role, error) {
			return nil, nil
		})
		//Test Scenario
		convey.Convey("Not enough information/id for binding the role object should return error", func() {
			patch := monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
				return "", errors.New("error")
			})
			c := &gin.Context{}
			GetRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Role is not exist in the database should return error", func() {
			patch := monkey.Patch(models.GetRoleByID, func(_ bson.ObjectId) (*models.Role, error) {
				return nil, errors.New("error")
			})
			c := &gin.Context{}
			GetRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Correct information has been passed to the db return error", func() {
			c := &gin.Context{}
			GetRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}

func TestDeleteRoleByID(t *testing.T) {
	convey.Convey("Testing DeleteRoleByID", t, func() {
		//Apply Patches
		monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
			return "id123", nil
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.Patch(models.DeleteRoleByID, func(_ bson.ObjectId) error {
			return nil
		})
		//Test Scenario
		convey.Convey("Not enough information/id for binding the role object should return error", func() {
			patch := monkey.Patch(common.StringToObjectID, func(_ string) (bson.ObjectId, error) {
				return "", errors.New("error")
			})
			c := &gin.Context{}
			DeleteRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Role is not exist in the database should return error", func() {
			patch := monkey.Patch(models.DeleteRoleByID, func(_ bson.ObjectId) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			DeleteRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Correct information has been passed to the db return error", func() {
			c := &gin.Context{}
			DeleteRoleByID(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})
	})
}
