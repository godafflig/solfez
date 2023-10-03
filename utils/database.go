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
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL, 
		email TEXT NOT NULL UNIQUE, 
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
		scoreLevel1 INTEGER DEFAULT 0 NOT NULL,
		scoreLevel2 INTEGER DEFAULT 0 NOT NULL,
		scoreLevel3 INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateScoreTableLevel1(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS scoreLevel1 (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		scoreLevel1 INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateScoreTableLevel2(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS scoreLevel2 (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		scoreLevel2 INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateScoreTableLevel3(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS scoreLevel3 (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		scoreLevel3 INTEGER DEFAULT 0 NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
