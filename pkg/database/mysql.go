package database

import (
	"database/sql"
	"github.com/spf13/viper"
)

func GetMySQL() (*sql.DB, error) {
	return sql.Open("mysql", viper.GetString("mysql"))
}
