// route for api module

package api

import (
	"sporule/api/app/modules/api/apicontrollers"
	"sporule/api/app/modules/middlewares"

	"github.com/gin-gonic/gin"
)

//RegisterAdminRoute register all routes for admin functions
func RegisterAdminRoute(router *gin.Engine) {

	apiRouter := router.Group("/admin/")
	//set route permission
	apiRouter.Use(middleware.JWTAuthMiddleware)

	apiRouter.GET("/user", apicontrollers.GetUsers)
}

//RegisterAuthRoute provides authentication functions such as generate token
func RegisterAuthRoute(router *gin.Engine) {
	authRouter := router.Group("/auth/")
	authRouter.POST("/user", apicontrollers.AddUser)
	authRouter.POST("/", apicontrollers.GenerateToken)

	//below are the auth router that requires token authentication
	requiredTokenRouter := authRouter.Group("/")
	requiredTokenRouter.Use(middleware.JWTAuthMiddleware)
	requiredTokenRouter.GET("/", apicontrollers.RefreshToken)
}

//RegisterTestRoute is for testing only
func RegisterTestRoute(router *gin.Engine) {
	testRouter := router.Group("/test/")
	testRouter.Use(middleware.RoleAuthMiddleware)
	testRouter.GET("/getusers", apicontrollers.GetUsers)
	testRouter.GET("/bbc", apicontrollers.GetUsers)
	testRouter.GET("/drop", apicontrollers.DropDB)
}
