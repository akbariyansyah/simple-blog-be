package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectionDB -> initilize connection to the database
func ConnectionDB() *sql.DB {

	DB_USER := ReadEnv("dbUser")

	DB_PASSWORD := ReadEnv("dbPassword")

	DB_HOST := ReadEnv("dbHost")

	DB_PORT := ReadEnv("dbPort")

	DB_SCHEMA := ReadEnv("dbName")

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA)
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Println(err)
	}
	log.Println("DB has started")
	return db
}
