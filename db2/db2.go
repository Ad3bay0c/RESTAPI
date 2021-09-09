package db2

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASS = "password"
	DB_NAME = "ApiTesting"
)
var DB *sql.DB

func init() {
	DB = connectDb()
	if DB.Ping() != nil {
		panic("Error Connecting to Database")
	}
	log.Printf("Database Connected")
}

func connectDb() *sql.DB{
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		log.Fatal("Error Connecting to Database")
		return nil
	}
	return db
}