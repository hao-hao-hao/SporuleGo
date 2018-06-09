package apicontrollers

import (
	"errors"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//AddUser provides the ability to add new user and return the added user with any error
func AddUser(c *gin.Context) {
	var tempUser models.User
	err := c.BindJSON(&tempUser)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
	}
	user, err := models.NewUser(tempUser.Email, tempUser.Password, tempUser.Name, tempUser.Roles)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
	}
	err = user.Register()
	if err == nil {
		err = errors.New("")
	} else {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
	}
	common.HTTPResponse200(c, &gin.H{"user": user}, common.GetError(err))
}

//GetUsers returns all the users
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(nil)
	common.HTTPResponse200(c, &gin.H{"users": users}, common.GetError(err))
}
