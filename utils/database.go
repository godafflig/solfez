package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// get the database
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

// create the 'users' table if it doesn't exist in the database
func CreateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL, 
		email TEXT NOT NULL UNIQUE, 
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_logged INTEGER DEFAULT 0 NOT NULL,
		score INTEGER DEFAULT 0 NOT NULL,
		profile_picture TEXT DEFAULT "https://visitemaroc.ca/wp-content/uploads/2021/06/profile-placeholder.png",
		wins INTEGER DEFAULT 0 NOT NULL,
		loses INTEGER DEFAULT 0 NOT NULL,
		total_games INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// create the 'scores' table if it doesn't exist in the database
func CreateScoreTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS scores (
		user_id INTEGER NOT NULL,
		user_name TEXT NOT NULL, 
		score INTEGER DEFAULT 0 NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (score) REFERENCES users(score)
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
