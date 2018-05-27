package common

import (
	"github.com/gin-gonic/gin"
)

//HTTPStatusStruct is the struct for http status
type HTTPStatusStruct struct {
	OK, MovedPermanently, BadRequest, Unauthorized, NotFound int
}

//LoadHTTPStatus sets the basic HTTPStatus
func (status *HTTPStatusStruct) LoadHTTPStatus() {
	status.OK = 200
	status.MovedPermanently = 301
	status.BadRequest = 400
	status.Unauthorized = 401
	status.NotFound = 404
}

//HTTPResponse is a generic response function which sets the structure of return
func HTTPResponse(c *gin.Context, code int, results *gin.H, err string) {
	body := gin.H{}
	if CheckNil(err) {
		body["error"] = err
	}
	if CheckNil(results) {
		body["results"] = *results
	}
	c.JSON(code, body)
}

//HTTPResponse404 returns 404 error
func HTTPResponse404(c *gin.Context) {
	//HTTPResponse(c, HTTPStatus.NotFound, nil, "Page Not Found")
	c.AbortWithStatus(HTTPStatus.NotFound)
}

//HTTPResponse401 returns 401 error
func HTTPResponse401(c *gin.Context) {
	//HTTPResponse(c, HTTPStatus.Unauthorized, nil, "You don't have necessary permission")
	c.AbortWithStatus(HTTPStatus.Unauthorized)
}

//HTTPResponse200 return 200 OK with results
func HTTPResponse200(c *gin.Context, results *gin.H, err string) {
	HTTPResponse(c, HTTPStatus.OK, results, err)
}
