package middleware

import (
	"errors"
	"net/http"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"
	"testing"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestJWTAuthMiddleware(t *testing.T) {
	convey.Convey("Testing JWTAuthMiddleware", t, func() {
		correctEmail := "abc@qq.com"
		//token, salt and email is all the same as correctEmail
		monkey.Patch(models.GetUserByEmail, func(email string) (*models.User, error) {
			if email == correctEmail {
				return &models.User{TokenSalt: email}, nil
			}
			return nil, errors.New("error")
		})
		monkey.Patch(common.HTTPResponse401, func(c *gin.Context) {
			c.Set("errors", "error")
			return
		})
		monkey.Patch(common.SetIDInHeader, func(c *gin.Context, id string) {
			c.Set("id", id)
		})
		monkey.Patch(common.VerifyToken, func(tokenString string) (string, string, error) {
			return tokenString, tokenString, nil
		})
		convey.Convey("Testing with correct token, Errors should be nil when the auth token is valid,also the id should be the same as the email address ", func() {
			c := &gin.Context{}
			c.Request = &http.Request{}
			c.Request.Header = map[string][]string{}
			c.Accepted = []string{}
			c.Request.Header.Set("Authorization", "bearer abc@qq.com")
			JWTAuthMiddleware(c)
			errors, _ := c.Get("errors")
			outputID, _ := c.Get("id")
			convey.So(errors, convey.ShouldBeNil)
			convey.So(outputID, convey.ShouldEqual, correctEmail)
		})
		convey.Convey("Testing with wrong token, Errors should not be empty, the id should be empty", func() {
			c := &gin.Context{}
			c.Request = &http.Request{}
			c.Request.Header = map[string][]string{}
			c.Accepted = []string{}
			c.Request.Header.Set("Authorization", "bearer abc@bb.com")
			JWTAuthMiddleware(c)
			errors, _ := c.Get("errors")
			outputID, _ := c.Get("id")
			convey.So(errors, convey.ShouldNotBeNil)
			convey.So(outputID, convey.ShouldBeNil)
		})
	})
}
