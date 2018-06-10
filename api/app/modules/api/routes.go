// route for api module

package api

import (
	"sporule/api/app/modules/api/apicontrollers"
	"sporule/api/app/modules/common"
	"sporule/api/app/modules/middlewares"

	"github.com/gin-gonic/gin"
)

//RegisterAPIRoutesV1 registers all api routers
func RegisterAPIRoutesV1(router *gin.Engine) {
	r := router.Group("/api/v1")
	registerAdminRoute(r)
	registerAuthRoute(r)
	registerFrontEndRoute(r)

	//test router, only activate when it is in "dev" environment
	registerTestRoute(r)
}

//registerAdminRoute register all routes for admin functions
func registerAdminRoute(router *gin.RouterGroup) {
	r := router.Group("/admin")
	//set route permission
	r.Use(middleware.JWTAuthMiddleware)
	r.GET("/users", apicontrollers.GetUsers)

	//fields
	r.GET("/fields", apicontrollers.GetFields)
	r.GET("/fields/:id", apicontrollers.GetFieldByID)
	r.POST("/fields", apicontrollers.AddField)
	r.PUT("/fields/:id", apicontrollers.UpdateField)
	r.DELETE("/fields/:id", apicontrollers.DeleteField)

	//roles
	r.GET("/roles", apicontrollers.GetRoles)
	r.GET("/roles/:id", apicontrollers.GetRoleByID)
	r.POST("/roles", apicontrollers.AddRole)
	r.PUT("/roles/:id", apicontrollers.UpdateRole)
	r.DELETE("/roles/:id", apicontrollers.DeleteRole)
}

//registerAuthRoute provides authentication functions such as generate token
func registerAuthRoute(router *gin.RouterGroup) {
	r := router.Group("/auth")
	r.POST("/user", apicontrollers.AddUser)
	r.POST("/", apicontrollers.GenerateToken)

	//below are the auth router that requires token authentication
	requiredTokenRouter := r.Group("/")
	requiredTokenRouter.Use(middleware.JWTAuthMiddleware)
	requiredTokenRouter.GET("/", apicontrollers.RefreshToken)
}

//registerFrontEndRoute is the routers that does not require any authentications
func registerFrontEndRoute(router *gin.RouterGroup) {
	r := router.Group("/")
	print(r)
}

//registerTestRoute is for testing only, it will only register the routes if it is in "dev" environment
func registerTestRoute(router *gin.RouterGroup) {
	//test routes
	if common.Config.ENV == "dev" {
		r := router.Group("/test/")
		r.Use(middleware.RoleAuthMiddleware)
		r.GET("/getusers", apicontrollers.GetUsers)
		r.GET("/drop", apicontrollers.DropDB)
	}

}
