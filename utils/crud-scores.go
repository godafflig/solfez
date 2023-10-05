package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func CreateScore(db *sql.DB, username string, userid int) {
	query := `
	INSERT INTO scores (user_id, user_name, score) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userid, username, "0")
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteScore(db *sql.DB, email string) {
	query := `
	DELETE FROM scores WHERE email = ?`
	_, err := db.Exec(query, email)
	if err != nil {
		log.Fatal(err)
	}
}

func getScore(db *sql.DB, email string) int {
	query := `
	SELECT score FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var score string
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			fmt.Println(err)
		}
	}

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		fmt.Println(err)
	}
	return scoreInt
}

func updateScore(db *sql.DB, email string, score int) {
	query := `
	UPDATE users SET score = ? WHERE email = ?`
	_, err := db.Exec(query, score, email)
	if err != nil {
		fmt.Println(err)
	}
}
