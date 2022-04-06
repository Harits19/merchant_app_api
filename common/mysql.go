package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitMysql() {

	var err error
	Db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/merchant_app")

	if err != nil {
		panic(err.Error())
	}
}
