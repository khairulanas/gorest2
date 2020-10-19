package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = `user04:1234567@/be04`

func ConnectMysql() (db *sql.DB, err error) {
	db, err = sql.Open(`mysql`,dsn)
	if err != nil {
		fmt.Println(err)
	}
	return
}
