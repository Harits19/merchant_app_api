package applicant

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
)

func ApplicantRouter(router *gin.RouterGroup) {
	router.POST("/", InsertApplicant)
	router.GET("/:name", DetailApplicant)
	router.DELETE("/:name", DeleteApplicant)
	router.PUT("/:name", UpdateApplicant)
}

func InsertApplicant(c *gin.Context) {
	validator := ApplicantModel{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err, _ := FindOneUser("email", validator.Email); err == nil {
		c.JSON(http.StatusForbidden, common.NewError("register", errors.New("this email already apply")))
		return
	}

	if err := InsertMongo(validator); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusCreated, validator)
}

func DetailApplicant(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("DetailApplicant : ", name)
	err, model := FindOneUser("name", name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

func DeleteApplicant(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("DetailApplicant : ", name)
	err := DeleteMongo(name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprint("Success delete applicant ", name))
}

func UpdateApplicant(c *gin.Context) {
	name := c.Param("name")
	validator := ApplicantModel{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err, _ := FindOneUser("name", name); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusForbidden, common.NewError("update", errors.New("name not found")))
		return
	}

	if err := UpdateMongo(validator); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, common.NewError("update", errors.New("error update")))
	}

	c.JSON(http.StatusOK, "success update")
}
