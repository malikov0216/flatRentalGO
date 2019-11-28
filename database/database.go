package database

import (
	"database/sql"
	"fmt"

	"github.com/malikov0216/createDataBase/config"
)

var DB *sql.DB

func Open() {
	var err error
	dbInfo := fmt.Sprintf("host=%s port=%s search_path=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbSchema, config.DbName, config.DbUser, config.DbPassword)

	DB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to PostgreSQL DB was established!")
}

func Close() error {
	return DB.Close()
}
