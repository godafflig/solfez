package utils

import (
	"database/sql"
	"log"
)

// rajouter le hash des mdp avant d'enregistrer dans bdd
func CreateUser(db *sql.DB, username string, password string, email string) {
	query := `
	INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err := db.Exec(query, username, password, email)
	if err != nil {
		log.Fatal(err)
	}
}

// vérifier que email entré = email bdd
// vérifier que hash mot de passe = hash mdp bdd
func CheckIfUserExist(db *sql.DB, email string, password string) bool {
	query := `
	SELECT email, password FROM users WHERE email = ? AND password = ?`
	rows, err := db.Query(query, email, password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}
