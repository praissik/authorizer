package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetMySQL() (*sql.DB, error) {
	return sql.Open("mysql", viper.GetString("mysql"))
}
