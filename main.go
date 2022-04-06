package main

import (
	"github.com/gin-gonic/gin"
	"jobapp.com/m/applicant"
	"jobapp.com/m/common"

	"net/http"
)

func main() {

	common.InitMongoDB()

	routerLocal := gin.Default()
	applicantGroup := routerLocal.Group("applicant")
	applicant.ApplicantRouter(applicantGroup)

	routerLocal.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routerLocal.Run()

}
