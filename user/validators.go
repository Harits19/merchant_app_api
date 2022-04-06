package user

import (
	"encoding/hex"

	"crypto/md5"

	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
)

const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"

func (user *UserModel) Bind(c *gin.Context) error {
	err := common.Bind(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (login *LoginModel) Bind(c *gin.Context) error {
	err := common.Bind(c, login)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) IsPasswordMatch(password string) bool {

	hash := md5.Sum([]byte(password))
	passwordLogin := hex.EncodeToString(hash[:])
	passwordDb := u.Password
	return passwordLogin == passwordDb
}
