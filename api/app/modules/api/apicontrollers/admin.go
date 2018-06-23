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
		return
	}
	user, err := models.NewUser(tempUser.Email, tempUser.Password, tempUser.Name, tempUser.Roles)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = user.Register()
	if err == nil {
		err = errors.New("")
	} else {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"user": user}, common.GetError(err))
}

//GetUsers returns all the users
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(nil)
	common.HTTPResponse200(c, &gin.H{"users": users}, common.GetError(err))
}

//GetFields returns all fields
func GetFields(c *gin.Context) {
	fields, err := models.GetFields(nil)
	common.HTTPResponse200(c, &gin.H{"fields": fields}, common.GetError(err))
}

//GetFieldByID return field by ID
func GetFieldByID(c *gin.Context) {
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	field, err := models.GetFieldByID(id)
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	common.HTTPResponse200(c, &gin.H{"field": field}, "")
}

//AddField provides the ability to add new field and return the added field with any error
func AddField(c *gin.Context) {
	var tempField models.Field
	err := c.BindJSON(&tempField)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	field, err := models.NewField(tempField.Name, tempField.FieldType)
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	err = field.Insert()
	if err != nil {
		common.HTTPResponse200(c, &gin.H{}, common.GetError(err))
		return
	}
	common.HTTPResponse200(c, &gin.H{"field": field}, "")
}

//DeleteField deletes the field
func DeleteField(c *gin.Context) {
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	err = models.DeleteField(id)
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	common.HTTPResponse200(c, &gin.H{}, "")
}

//UpdateField updates the field
func UpdateField(c *gin.Context) {
	var field models.Field
	id, err := common.StringToObjectID(c.Param("id"))
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	err = c.BindJSON(&field)
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	err = field.Update(id)
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	common.HTTPResponse200(c, &gin.H{"field": field}, "")
}

//AddRole adds new role to the system db
func AddRole(c *gin.Context) {
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
