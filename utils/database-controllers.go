package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// rajouter le hash des mdp avant d'enregistrer dans bdd
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

// vérifier que hash mot de passe = hash mdp bdd
func userExists(db *sql.DB, email string, password string) bool {
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

	fmt.Println("storedPassword : ", storedPassword)
	fmt.Println("password : ", password)

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func usernameExists(db *sql.DB, username string) bool {
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

func emailExists(db *sql.DB, email string) bool {
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

func getId(db *sql.DB, email string) int {
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
func getUsername(db *sql.DB, email string) string {
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
