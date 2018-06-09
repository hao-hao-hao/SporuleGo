//main.go

package main

import (
	"sporule/api/app/modules/api"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

func main() {

	//initiate application
	common.InitiateGlobalVariables()
	common.InitiateDatabase()

	//set router to gin default router with logger
	router := gin.Default()

	//Register the routes
	api.RegisterAPIRoutesV1(router)

	//start the application
	router.Run()
}
