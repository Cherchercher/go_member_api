package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var db *sql.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn, err := sql.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func GetDB() *sql.DB {
	return db
}
