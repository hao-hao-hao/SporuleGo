//main.go

package main

import (
	"sporule/app/modules/frontend"
	"ui/app/modules/common"

	"github.com/gin-gonic/gin"
)

func main() {
	//set router to gin default router with logger
	router := gin.Default()

	//Load the template
	//router.LoadHTMLGlob("templates/*/*.html")

	//Initiate the Database
	common.InitiateDatabase()

	//Initiate the routes
	frontend.RegisterFrontEndRoute(router)
	//start the application
	router.Run()
}
