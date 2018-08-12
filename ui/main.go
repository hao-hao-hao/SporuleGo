//main.go

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//set router to gin default router with logger
	router := gin.Default()

	//Load the templates
	//router.LoadHTMLGlob("templates/*/*.html")

	//Initiate the routes

	//start the application
	router.Run()
}
