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
)

func TestGenerateToken(t *testing.T) {
	convey.Convey("Testing GenerateToken", t, func() {
		//Apply Patches
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(_ *gin.Context, _ interface{}) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "Verify", func(_ *models.User) error {
			return nil
		})
		monkey.Patch(common.GenerateJWT, func(_, _ string) (string, error) {
			return "tokenstring", nil
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		convey.Convey("User credential binding failure should return error", func() {
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "BindJSON", func(_ *gin.Context, _ interface{}) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			GenerateToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("User credential Verify failure should return error", func() {
			patch := monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "Verify", func(_ *models.User) error {
				return errors.New("error")
			})
			c := &gin.Context{}
			GenerateToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("User credential passed all the checks should return tokenstring", func() {
			c := &gin.Context{}
			GenerateToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})

	})
}

func TestRefreshToken(t *testing.T) {
	convey.Convey("Testing RefreshToken", t, func() {
		//Apply Patches
		monkey.Patch(common.GetIDInHeader, func(_ *gin.Context) string {
			return "abc@qq.com"
		})
		monkey.Patch(common.HTTPResponse401, func(c *gin.Context) {
			c.Set("results", "error")
		})
		monkey.Patch(models.GetUserByEmail, func(_ string) (*models.User, error) {
			return &models.User{Name: "hi"}, nil
		})
		monkey.Patch(common.HTTPResponse200, func(c *gin.Context, results *gin.H, err string) {
			if len(*results) > 0 {
				c.Set("results", "result")
			} else {
				c.Set("results", "error")
			}
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(&models.User{}), "UpdateTokenSalt", func(_ *models.User) error {
			return nil
		})
		monkey.Patch(common.GenerateJWT, func(_, _ string) (string, error) {
			return "tokenstring", nil
		})
		convey.Convey("UnAuthorised(No ID in header) will return error", func() {
			patch := monkey.Patch(common.GetIDInHeader, func(_ *gin.Context) string {
				return ""
			})
			c := &gin.Context{}
			RefreshToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Invalid header will return error", func() {
			patch := monkey.Patch(models.GetUserByEmail, func(_ string) (*models.User, error) {
				return nil, errors.New("error")
			})
			c := &gin.Context{}
			RefreshToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "error")
			patch.Unpatch()
		})
		convey.Convey("Authroised user will return the token", func() {
			c := &gin.Context{}
			RefreshToken(c)
			results, _ := c.Get("results")
			convey.So(results, convey.ShouldEqual, "result")
		})

	})
}
