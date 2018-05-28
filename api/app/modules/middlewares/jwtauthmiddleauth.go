package middleware

import (
	"sporule/api/app/models"
	"sporule/api/app/modules/common"
	"strings"

	"github.com/gin-gonic/gin"
)

//JWTAuthMiddleware is the middleware that authenticates the requests
func JWTAuthMiddleware(c *gin.Context) {
	//need to add role authroisation

	//before request
	authHeader := c.Request.Header.Get(common.Config.AuthHeader)
	if common.CheckNil(authHeader) {
		authToken := strings.SplitN(authHeader, " ", 2)[1]
		email, _ := common.VerifyToken(authToken)
		if common.CheckNil(email) {
			user, _ := models.GetUserByEmail(email)
			if common.CheckNil(user) {
				c.Set("email", email)
			} else {
				common.HTTPResponse401(c)
			}

		} else {
			common.HTTPResponse401(c)
		}
	} else {
		common.HTTPResponse401(c)
	}

	c.Next()

	//after request
}
