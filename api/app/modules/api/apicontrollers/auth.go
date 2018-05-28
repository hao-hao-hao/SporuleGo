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
			tokenString, _ := common.GenerateJWT(user.Email)
			common.HTTPResponse200(c, &gin.H{"Token": tokenString}, common.GetError(err))
		}
	}
	common.HTTPResponse200(c, nil, common.GetError(err))
}
