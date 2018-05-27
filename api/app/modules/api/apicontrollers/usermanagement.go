package apicontrollers

import (
	"errors"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//AddUser provides the ability to add new user and return the added user with any error
func AddUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		err = user.Register()
	}
	if err == nil {
		err = errors.New("No Error")
	}
	common.HTTPResponse200(c, &gin.H{"user": user}, err.Error())
}

//GetUsers returns all the users
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(nil)
	common.HTTPResponse200(c, &gin.H{"users": users}, err.Error())
}
