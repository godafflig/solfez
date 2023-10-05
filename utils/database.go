package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	var db *sql.DB

	if _, err := os.Stat("./database/"); os.IsNotExist(err) {
		err = nil
		db, err = sql.Open("sqlite3", "./database/database.db")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db, err = sql.Open("sqlite3", "./database/database.db")
		if err != nil {
			log.Println("./database/database.db")
			log.Fatal(err.Error() + "./database/database.db")
		}
	}
	return db
}

func CreateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL, 
		email TEXT NOT NULL UNIQUE, 
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_logged INTEGER DEFAULT 0 NOT NULL,
		scoreLevel1 INTEGER DEFAULT 0 NOT NULL,
		scoreLevel2 INTEGER DEFAULT 0 NOT NULL,
		scoreLevel3 INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateScoreTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS scores (
		user_id INTEGER NOT NULL, 
		score INTEGER DEFAULT 0 NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (score) REFERENCES users(score)
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
