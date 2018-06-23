package apicontrollers

import (
	"sporule/api/app/models"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//GenerateToken verifies the user credentials and then returns a jwt
func GenerateToken(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		err = user.Verify()
		if err == nil {
			tokenString, _ := common.GenerateJWT(user.Email, user.TokenSalt)
			common.HTTPResponse200(c, &gin.H{"Token": tokenString}, common.GetError(err))
			return
		}
	}
	common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
}

//RefreshToken expires the invalidate existing token and provide a new token
func RefreshToken(c *gin.Context) {
	email := common.GetIDInHeader(c)
	if !common.CheckNil(email) {
		common.HTTPResponse401(c)
		return
	}
	user, err := models.GetUserByEmail(email)
	if common.CheckNil(err) {
		common.HTTPResponse401(c)
		return
	}
	if user.UpdateTokenSalt() != nil {
		common.HTTPResponse200(c, &gin.H{}, common.Enums.ErrorMessages.SystemError)
		return
	}
	tokenString, _ := common.GenerateJWT(user.Email, user.TokenSalt)
	common.HTTPResponse200(c, &gin.H{"Token": tokenString}, common.GetError(err))
}
