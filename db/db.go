package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlCreds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlCreds)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_, err = db.Exec(`
		create table if not exists refresh_tokens (
			user_id text primary key,
			token_hash text not null 
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveRefreshToken(userID, tokenHash string) error {
	_, err := db.Exec("insert into refresh_tokens (user_id, token_hash) values ($1, $2)", userID, tokenHash)
	return err
}

func GetRefreshToken(userID string) (string, error) {
	var tokenHash string
	err := db.QueryRow("select token_hash from refresh_tokens where user_id = $1", userID).Scan(&tokenHash)
	if err != nil {
		return "", err
	}
	return tokenHash, nil
}

func UpdateRefreshToken(userID, tokenHash string) error {
	_, err := db.Exec("update refresh_tokens set token_hash = $1 where user_id = $2", tokenHash, userID)
	return err
}
