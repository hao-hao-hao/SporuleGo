//main.go

package main

import (
	"sporule/api/app/modules/api"
	"sporule/api/app/modules/common"

	"github.com/gin-gonic/gin"
)

func main() {
	//initiate Global Variables
	common.InitiateGlobalVariables()

	//Initiate the Database
	common.InitiateDatabase()
	//set router to gin default router with logger
	router := gin.Default()

	//Initiate the routes
	api.RegisterAPIRoute(router)

	//start the application
	router.Run()
}
