package utils

import (
	"database/sql"
	"fmt"
	"log"
)

// rajouter le hash des mdp avant d'enregistrer dans bdd
func CreateUser(db *sql.DB, username string, password string, email string) {
	query := `
	INSERT INTO users (username, password, email, is_logged) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, username, password, email, "1")
	if err != nil {
		log.Fatal(err)
	}
}

// vérifier que hash mot de passe = hash mdp bdd
func userExists(db *sql.DB, email string, password string) bool {
	query := `
	SELECT email, password FROM users WHERE email = ? AND password = ?`
	rows, err := db.Query(query, email, password)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
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
