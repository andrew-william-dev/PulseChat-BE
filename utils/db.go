package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:EdHub@localhost:5432/chat-app?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
	log.Println("Connected to PostgreSQL")
}
