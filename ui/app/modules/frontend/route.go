// route for front_end module

package frontend

import (
	"ui/app/modules/frontend/frontendcontroller"

	"github.com/gin-gonic/gin"
)

var frontEndRouter *gin.RouterGroup

//RegisterFrontEndRoute register all routes for front end
func RegisterFrontEndRoute(router *gin.Engine) {

	frontEndRouter = router.Group("/")
	frontEndRouter.GET("/", frontendcontroller.Index)
}
