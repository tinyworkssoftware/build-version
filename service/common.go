package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

func connectDb() (*sqlx.DB, error) {
	return sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_URL"),
			os.Getenv("DB_NAME"),
		))
}