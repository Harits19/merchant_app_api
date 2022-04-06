package applicant

import (
	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
)

func (self *ApplicantModel) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	return nil
}
