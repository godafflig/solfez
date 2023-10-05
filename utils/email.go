package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
)

var emailVerificationCache = cache.New(5*time.Minute, 10*time.Minute)

func SaveEmailemailVerificationCache(token string) {
	emailVerificationCache.Set(token, nil, cache.DefaultExpiration)
}

func GenerateEmailVerificationToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}

func SendEmailConfirmation(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "solfez@gmx.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Confirmation d'email")
	m.SetBody("text/html", "Cliquez sur ce lien pour confirmer votre email : <a href='http://localhost:3000/login'>Confirmer</a>")

	d := gomail.NewDialer("mail.gmx.com", 587, "solfez@gmx.com", "kKg4FNs3xMNPbpm6")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// Si le jeton de confirmation d'e-mail est valide.
func IsValidEmailVerificationToken(token string) bool {
	// Recherche jeton bdd
	query := `
        SELECT user_id
        FROM email_verification_tokens
        WHERE token = ?`
	var userId int
	err := GetDB().QueryRow(query, token).Scan(&userId)
	if err != nil {
		// Le jeton n'a pas été trouvé dans la bdd
		fmt.Println("Invalid email verification token:", token)
		return false
	}

	return true
}

// Activation du compte de l'utilisateur correspondant au jeton de confirmation.
func ActivateUserAccount(token string) {
	// Recherchez l'utilisateur correspondant au jeton dans la base de données
	query := `
        SELECT user_id
        FROM email_verification_tokens
        WHERE token = ?`
	var userId int
	err := GetDB().QueryRow(query, token).Scan(&userId)
	if err != nil {
		fmt.Println("Error finding user for activation:", err)
		return
	}

	// Mettez à jour le statut d'activation de l'utilisateur dans la base de données
	updateQuery := `
        UPDATE users
        SET is_active = 1
        WHERE user_id = ?`
	_, err = GetDB().Exec(updateQuery, userId)
	if err != nil {
		fmt.Println("Error activating user account:", err)
	}
}

func EmailConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if tempUser, ok := tempUserStore.Load(token); ok {
		if tempUser, ok := tempUser.(*TempUser); ok {
			tempUser.Confirmed = true
			CreateUser(GetDB(), tempUser.Username, tempUser.Password, tempUser.Email)
			tempUserStore.Delete(token)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/confirmation-error", http.StatusSeeOther)
}
