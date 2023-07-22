package config

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/utils"
)

var (
	dbHost     = utils.DotEnv("DB_HOST")
	dbPort     = utils.DotEnv("DB_PORT")
	dbUser     = utils.DotEnv("DB_USER")
	dbPassword = utils.DotEnv("DB_PASSWORD")
	dbName     = utils.DotEnv("DB_NAME")
	sslMode    = utils.DotEnv("SSL_MODE")
)

// db connection
var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

func ConnectDB() (*sql.DB) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connected")
	}
	return db
}
