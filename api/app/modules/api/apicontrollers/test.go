package apicontrollers

import (
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

//DropDB drops the database
func DropDB(c *gin.Context) {
	common.DropDatabase()
	common.HTTPResponse200(c, &gin.H{}, "")
}
