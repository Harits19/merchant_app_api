package user

import (
	"fmt"

	"jobapp.com/m/common"
)

type UserModel struct {
	Id       string               `json:"id"`
	Name     string               `json:"name" binding:"required"`
	UserName string               `json:"user_name" binding:"required"`
	Password string               `json:"password,omitempty" binding:"required"`
	Modified common.ModifiedModel `json:"modified"`
}

type LoginModel struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FindOneSql(userName string) (UserModel, error) {

	var model UserModel

	row := common.Db.QueryRow("SELECT id, name, user_name, password, created_at, created_by, updated_at, updated_by FROM users WHERE user_name=?", userName)

	err := row.Scan(&model.Id, &model.Name, &model.UserName, &model.Password, &model.Modified.CreatedAt, &model.Modified.CreatedBy, &model.Modified.UpdatedAt, &model.Modified.UpdatedBy)
	if err != nil {
		return model, err
	}
	fmt.Println("result", model)

	return model, nil
}
