package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASS = "password"
	DB_NAME = "ApiTesting"
)

func ConnectDb() *sql.DB{
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		log.Fatal("Error Connecting to Database")
		return nil
	}
	return db
}