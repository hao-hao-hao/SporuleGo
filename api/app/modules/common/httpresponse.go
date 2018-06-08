package common

import (
	"github.com/gin-gonic/gin"
)

//HTTPStatusStruct is the struct for http status
type hTTPStatusStruct struct {
	OK, MovedPermanently, BadRequest, Unauthorized, NotFound int
}

//LoadHTTPStatus sets the basic HTTPStatus
func (enums *enum) loadHTTPStatus() {
	enums.HTTPStatus.OK = 200
	enums.HTTPStatus.MovedPermanently = 301
	enums.HTTPStatus.BadRequest = 400
	enums.HTTPStatus.Unauthorized = 401
	enums.HTTPStatus.NotFound = 404
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
	c.AbortWithStatus(Enums.HTTPStatus.NotFound)
}

//HTTPResponse401 returns 401 error
func HTTPResponse401(c *gin.Context) {
	//HTTPResponse(c, HTTPStatus.Unauthorized, nil, "You don't have necessary permission")
	c.AbortWithStatus(Enums.HTTPStatus.Unauthorized)
}

//HTTPResponse200 return 200 OK with results
func HTTPResponse200(c *gin.Context, results *gin.H, err string) {
	HTTPResponse(c, Enums.HTTPStatus.OK, results, err)
}
