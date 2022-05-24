package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	if DB != nil {
		err = DB.Ping()
		if err == nil {
			return
		}
	}

	DB, err = sql.Open("mysql", "root:root@/app")
	if err != nil {
		panic(err)
	}
}
