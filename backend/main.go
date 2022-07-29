package main

import (
	"backend/models"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "backend/db/seeds"

	"github.com/joho/godotenv"
	"github.com/kristijorgji/goseeder"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	models.Config()
	goseeder.WithSeeder(connectToDbOrDie, func() {})
}

func connectToDbOrDie() *sql.DB {
	dbDriver := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		url.QueryEscape(dbPassword),
		dbHost,
		dbPort,
		dbName,
	)
	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	return con
}
