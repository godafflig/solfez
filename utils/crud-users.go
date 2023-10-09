package utils

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// create a new user in database
func CreateUser(db *sql.DB, username string, password string, email string) {
	//hash des mdp
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	query := `
	INSERT INTO users (username, password, email, is_logged) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, username, hashedPassword, email, "1")
	if err != nil {
		log.Fatal(err)
	}
}

// delete one user from both database
func DeleteUser(db *sql.DB, email string, w http.ResponseWriter, r *http.Request) {
	query := `
	DELETE FROM users WHERE email = ?`
	_, err := db.Exec(query, email)
	if err != nil {
		fmt.Println(err)
	}

	query2 := `
	DELETE FROM scores WHERE user_id = ?`
	_, err = db.Exec(query2, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}

	SessionData.Id = 0
	SessionData.Username = ""
	SessionData.Email = ""
	SessionData.IsLogged = false
	SessionData.Error = ""
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.Score = 0
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// check if one user exists in the database
func UserExists(db *sql.DB, email string, password string) bool {
	query := `
	SELECT password FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var storedPassword string
	for rows.Next() {
		err := rows.Scan(&storedPassword)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// check if a username is already existing in the database
func UsernameExists(db *sql.DB, username string) bool {
	query := `
	SELECT username FROM users WHERE username = ?`
	rows, err := db.Query(query, username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

// check if an email is already existing in the database
func EmailExists(db *sql.DB, email string) bool {
	query := `
	SELECT email FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

// get one user id from the database based on the email
func GetId(db *sql.DB, email string) int {
	query := `
	SELECT user_id FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	return idInt
}

// get one username from the database based on the email
func GetUsername(db *sql.DB, email string) string {
	query := `
	SELECT username FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
		}
	}
	return username
}

// get one score from the database 'users' based on the email
func GetScore(db *sql.DB, email string) int {
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

// update one score from the database 'users' based on the email
func UpdateScore(db *sql.DB, email string, score int) {
	query := `
	UPDATE users SET score = ? WHERE email = ?`
	_, err := db.Exec(query, score, email)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUserPassword(db *sql.DB, email string, password string) {
	query := `
	UPDATE users SET password = ? WHERE email = ?`
	_, err := db.Exec(query, password, email)
	if err != nil {
		fmt.Println(err)
	}
}
