package utils

import (
	"fmt"
	"net/http"
	"sync"
)

type TempUser struct {
	Username  string
	Email     string
	Password  string
	Token     string
	Confirmed bool
}

var tempUserStore sync.Map

func Login(email string, password string, w http.ResponseWriter, r *http.Request) {
	if !emailExists(GetDB(), email) {
		SessionData.Error = "Wrong email."
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if !userExists(GetDB(), email, password) {
		SessionData.Error = "Wrong password."
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		SessionData.Id = getId(GetDB(), email)
		SessionData.Username = getUsername(GetDB(), email)
		SessionData.Email = email
		SessionData.IsLogged = true
		SessionData.Error = ""
		fmt.Println("Logged in:", SessionData)
		http.Redirect(w, r, "/play", http.StatusSeeOther)
	}
}

func Register(username string, email string, password string, passwordCheck string, w http.ResponseWriter, r *http.Request) {
	if password != passwordCheck {
		SessionData.Error = "Passwords don't match."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else if usernameExists(GetDB(), username) {
		SessionData.Error = "Username already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else if emailExists(GetDB(), email) {
		SessionData.Error = "Email already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Générer un jeton de confirmation d'e-mail
	token, err := GenerateEmailVerificationToken()
	if err != nil {
		// Gérer les erreurs de génération de jeton
		fmt.Println(err)
		http.Redirect(w, r, "/error-page", http.StatusSeeOther)
		return
	}

	tempUser := TempUser{
		Username:  username,
		Email:     email,
		Password:  password,
		Token:     token,
		Confirmed: false,
	}

	tempUserStore.Store(token, tempUser)
	err = SendEmailConfirmation(email, token)
	if err != nil {
		http.Redirect(w, r, "/error-page", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/test?token="+token, http.StatusSeeOther)
}

func RedirectTest(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Change is_logged to 0 in the database
	query := `
	UPDATE users SET is_logged = ? WHERE email = ?`
	_, err := GetDB().Exec(query, "0", SessionData.Email)
	if err != nil {
		panic(err)
	}

	fmt.Print("Confirmed logout: ", SessionData, " -> ")
	SessionData.Id = 0
	SessionData.Username = ""
	SessionData.Email = ""
	SessionData.IsLogged = false
	SessionData.Error = ""
	fmt.Println(SessionData)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func isStrongPassword(password string) bool {
	const (
		minLength    = 8
		minDigits    = 1
		minSymbols   = 1
		minUppercase = 1
	)

	if len(password) < minLength {
		return false
	}

	var minDigitsCount, minSymbolsCount, minUppercaseCount int
	for _, c := range password {
		switch {
		case '0' <= c && c <= '9':
			minDigitsCount++
		case 'a' <= c && c <= 'z':
		case 'A' <= c && c <= 'Z':
			minUppercaseCount++
		default:
			minSymbolsCount++
		}
	}

	if minDigitsCount < minDigits || minSymbolsCount < minSymbols || minUppercaseCount < minUppercase {
		return false
	}

	return true
}
