package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jobapp.com/m/user"
)

func Routers(router *gin.RouterGroup) {
	router.GET("/omzet", OmzetMerchant)
	router.GET("/omzet-outlet", OmzetMerchantOutlet)
}

func OmzetMerchant(c *gin.Context) {
	yearMonth := c.Query("year-month")
	modelJwt := c.MustGet("my_user_model").(user.UserModel)

	result, err := OmzetPerDay(modelJwt.Id, yearMonth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	serialize := OmzetWithoutOutlets(result)
	c.JSON(http.StatusOK, serialize)

}

func OmzetMerchantOutlet(c *gin.Context) {
	yearMonth := c.Query("year-month")
	modelJwt := c.MustGet("my_user_model").(user.UserModel)

	result, err := OmzetPerDay(modelJwt.Id, yearMonth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)

}
