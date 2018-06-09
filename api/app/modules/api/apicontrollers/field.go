package apicontrollers

import (
	"sporule/api/app/models"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//GetFields returns all fields
func GetFields(c *gin.Context) {
	fields, err := models.GetFields(nil)
	common.HTTPResponse200(c, &gin.H{"fields": fields}, common.GetError(err))
}

//GetFieldsByID return fields by ID
func GetFieldsByID(c *gin.Context) {
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
		common.HTTPResponse404(c)
		return
	}
	field, err := models.NewField(tempField.Name, tempField.FieldType)
	if err != nil {
		common.HTTPResponse404(c)
		return
	}
	err = field.Insert()
	if err != nil {
		common.HTTPResponse409(c)
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
	common.HTTPResponse200(c, &gin.H{}, common.Enums.ErrorMessages.PageNotFound)
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
	}
	err = field.Update(id)
	if err != nil {
		common.HTTPResponse404(c)
	}
	common.HTTPResponse200(c, &gin.H{"field": field}, "")
}
