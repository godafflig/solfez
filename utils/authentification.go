package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

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
		SessionData.Score = getScore(GetDB(), email)
		SessionData.Error = ""
		SessionData.ProfilePic = GetProfilePicFromDb()
		fmt.Println("Logged in : ", SessionData)

		StartGame(w, r)
		template.Must(template.ParseFiles("static/play.html")).Execute(w, SessionData)
	}
}

func Register(username string, email string, password string, passwordCheck string, w http.ResponseWriter, r *http.Request) {
	if password != passwordCheck {
		SessionData.Error = "Passwords don't match."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		// } else if !isStrongPassword(password) {
		// 	SessionData.Error = "Password must be at least 8 characters long and contain at least 1 digit, 1 symbol and 1 uppercase letter."
		// 	http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else if usernameExists(GetDB(), username) {
		SessionData.Error = "Username already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else if emailExists(GetDB(), email) {
		SessionData.Error = "Email already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else {
		CreateUser(GetDB(), username, password, email)
		SessionData.Id = getId(GetDB(), email)
		SessionData.Username = username
		SessionData.Email = email
		SessionData.IsLogged = true
		SessionData.Error = ""
		SessionData.ProfilePic = GetProfilePicFromDb()
		fmt.Println("Registered : ", SessionData)

		StartGame(w, r)
		template.Must(template.ParseFiles("static/play.html")).Execute(w, SessionData)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// change is_logged to 0 in database
	query := `
	UPDATE users SET is_logged = ? WHERE email = ?`
	_, err := GetDB().Exec(query, "0", SessionData.Email)
	if err != nil {
		panic(err)
	}

	fmt.Print("Confirmed logout : ", SessionData, " -> ")
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
