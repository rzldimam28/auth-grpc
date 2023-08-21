package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Connect() *sql.DB {

	dbUser := viper.GetString("db_user")
	dbPassword := viper.GetString("db_password")
	dbHost := viper.GetString("db_host")
	dbPort := viper.GetInt("db_port")
	dbName := viper.GetString("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("fatal connect to database: %w", err))
	}

	return db
}