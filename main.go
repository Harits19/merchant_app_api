package main

import (
	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
	"jobapp.com/m/merchant"
	"jobapp.com/m/transaction"
	"jobapp.com/m/user"

	"net/http"
)

func main() {

	common.InitMysql()
	routerLocal := gin.Default()
	user.Routers(routerLocal.Group("user"))
	merchant.Routers(routerLocal.Group("merchant"))
	routerLocal.Use(user.AuthMiddleware(true))
	transaction.Routers(routerLocal.Group("transaction"))

	routerLocal.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routerLocal.Run()

}
