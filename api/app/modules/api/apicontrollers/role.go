package apicontrollers

import (
	"fmt"
	"sporule/api/app/models"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//AddRole adds new role to the system db
func AddRole(c *gin.Context) {
	fmt.Printf("%v", c.Request.Body)
	var tempRole models.Role
	err := c.BindJSON(&tempRole)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	role, err := models.NewRole(tempRole.Name)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = role.Insert()
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"role": role}, "")
}

//UpdateRole updates the role by id
func UpdateRole(c *gin.Context) {
	var role models.Role
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = c.BindJSON(&role)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = role.Update(id)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"role": role}, "")
}

//GetRoles returns all fields
func GetRoles(c *gin.Context) {
	roles, err := models.GetRoles(nil)
	common.HTTPResponse200(c, &gin.H{"roles": roles}, common.GetError(err))
}

//GetRoleByID return role by ID
func GetRoleByID(c *gin.Context) {
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	role, err := models.GetRoleByID(id)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"role": role}, "")
}

//DeleteRoleByID deletes the role by id
func DeleteRoleByID(c *gin.Context) {
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = models.DeleteRoleByID(id)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"result": "Deleted"}, "")
}
