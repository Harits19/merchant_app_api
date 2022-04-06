package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
)

func Routers(r *gin.RouterGroup) {

	r.POST("/", UserLogin)
}

func UserLogin(c *gin.Context) {

	login := LoginModel{}

	if err := login.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	model, err := FindOneSql(login.UserName)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusForbidden, "username not found")
		return
	}

	if !model.IsPasswordMatch(login.Password) {
		c.JSON(http.StatusForbidden, "invalid password")
		return
	}

	token := GenToken(login.UserName)

	c.JSON(http.StatusCreated, gin.H{
		"user":  model,
		"token": token,
	})

}
