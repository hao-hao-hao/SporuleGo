// route for api module

package api

import (
	"sporule/api/app/modules/api/apicontrollers"
	"sporule/api/app/modules/middleware"

	"github.com/gin-gonic/gin"
)

var apiRouter *gin.RouterGroup

//RegisterAPIRoute register all routes for API
func RegisterAPIRoute(router *gin.Engine) {

	apiRouter = router.Group("/")
	//set route permission
	apiRouter.Use(middleware.JWTAuthMiddleware)

	apiRouter.POST("/user", apicontrollers.AddUser)
	apiRouter.GET("/", apicontrollers.GetUsers)
}
