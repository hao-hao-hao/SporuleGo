package common

import (
	"github.com/gin-gonic/gin"
)

//HTTPResponse is a generic response function which sets the structure of return
func HTTPResponse(c *gin.Context, code int, results *gin.H, err string) {
	body := gin.H{}
	if CheckNil(err) {
		body["errors"] = err
	} else if CheckNil(results) {
		body["data"] = *results
	}
	c.JSON(code, body)
}

//HTTPResponse404 returns 404 error
func HTTPResponse404(c *gin.Context) {
	//HTTPResponse(c, HTTPStatus.NotFound, nil, "Page Not Found")
	c.AbortWithStatus(Enums.HTTPStatus.NotFound)
}

//HTTPResponse401 returns 401 error
func HTTPResponse401(c *gin.Context) {
	//HTTPResponse(c, HTTPStatus.Unauthorized, nil, "You don't have necessary permission")
	c.AbortWithStatus(Enums.HTTPStatus.Unauthorized)
}

//HTTPResponse200 return 200 OK with results
func HTTPResponse200(c *gin.Context, results *gin.H, err string) {
	if results == nil {
		results = &gin.H{}
	}
	HTTPResponse(c, Enums.HTTPStatus.OK, results, err)
}

//HTTPResponse409 is the conflict response
func HTTPResponse409(c *gin.Context) {
	c.AbortWithStatus(Enums.HTTPStatus.Conflict)
}

//HTTPResponse204 is the no content response
func HTTPResponse204(c *gin.Context) {
	c.AbortWithStatus(Enums.HTTPStatus.NoContent)
}
