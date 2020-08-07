package repository

import (
	model "build-version/model/toml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDb(conf *model.TomlConfig) (*sqlx.DB, error) {
	return sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s",
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.DatabaseUrl,
			conf.Database.DatabaseName,
		))
}