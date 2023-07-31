package db

import (
	"database/sql"
	"fmt"
	"os"
)

var psDb *sql.DB

func NewPostgresStore() error {

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbName == "" || dbUser == "" || dbPassword == "" {
		return fmt.Errorf("database credentials missing")
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbName, dbUser, dbPassword)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	psDb = &sql.DB{}

	return nil
}

func GetDB() *sql.DB {
	return psDb
}
