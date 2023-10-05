package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

// create one score (related the a user id & user name) in database
func CreateScore(db *sql.DB, username string, userid int) {
	query := `
	INSERT INTO scores (user_id, user_name, score) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userid, username, "0")
	if err != nil {
		log.Fatal(err)
	}

}

// delete one score from the database
func DeleteScore(db *sql.DB, email string) {
	query := `
	DELETE FROM scores WHERE email = ?`
	_, err := db.Exec(query, email)
	if err != nil {
		log.Fatal(err)
	}
}

// get score from score bdd
func getScoreFromScoresTable() int {
	db := GetDB()
	query := `
	SELECT score FROM scores WHERE user_id = ?`
	rows, err := db.Query(query, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var oldScoreString string
	for rows.Next() {
		err := rows.Scan(&oldScoreString)
		if err != nil {
			fmt.Println(err)
		}
	}

	oldScore, err := strconv.Atoi(oldScoreString)
	if err != nil {
		fmt.Println(err)
	}
	return oldScore
}

// update one score in the database
func updateScoreInScoreTable(db *sql.DB, id int, newScore int) {
	query := `
		UPDATE scores SET score = ? WHERE user_id = ?`
	_, err := db.Exec(query, newScore, id)
	if err != nil {
		fmt.Println(err)
	}
}
