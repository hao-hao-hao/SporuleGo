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
	//Locate Authorization Header
	authHeader := c.Request.Header.Get("Authorization")
	if !common.CheckNil(authHeader) {
		common.HTTPResponse401(c)
	}
	//Check if it is bearer token
	authString := strings.SplitN(authHeader, " ", 2)
	if strings.ToLower(authString[0]) != "bearer" {
		common.HTTPResponse401(c)
	}
	//Check if the token is valid
	authToken := authString[1]
	email, salt, _ := common.VerifyToken(authToken)
	if !common.CheckNil(email, salt) {
		common.HTTPResponse401(c)
	}
	//Check if the user is a valid user
	user, err := models.GetUserByEmail(email)
	if common.CheckNil(err) || user.TokenSalt != salt {
		common.HTTPResponse401(c)
	}
	//attach email address in the header
	common.SetIDInHeader(c, email)

	c.Next()

	//after request
}
