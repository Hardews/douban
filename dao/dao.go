package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:lmh123@tcp(127.0.0.1:3306)/JD_database")
	if err != nil {
		fmt.Println(err)
	}

	dB = db
}
