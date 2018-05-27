package frontendcontroller

import (
	"net/http"
	"sporule/ui/app/model"
	"ui/app/model"

	"github.com/gin-gonic/gin"
)

//Index Controller for Home
func Index(c *gin.Context) {
	user := model.NewUser("haleluohao@gmail.com", "1q2w3e4r", "Hao", nil)
	user.Register()
	users, err := model.GetUserByEmail()
	print(users, err)
	c.String(http.StatusOK, "Hello World")
}
