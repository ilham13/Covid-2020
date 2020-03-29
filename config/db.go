package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/covid-2020")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
