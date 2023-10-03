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
		SessionData.Error = ""
		fmt.Println("Logged in : ", SessionData)
		template.Must(template.ParseFiles("static/play.html")).Execute(w, SessionData)
	}
}

func Register(username string, email string, password string, passwordCheck string, w http.ResponseWriter, r *http.Request) {
	if password != passwordCheck {
		SessionData.Error = "Passwords don't match."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else if usernameExists(GetDB(), username) {
		SessionData.Error = "Username already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else if emailExists(GetDB(), email) {
		SessionData.Error = "Email already exists."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	} else {
		SessionData.Id = getId(GetDB(), email)
		SessionData.Username = username
		SessionData.Email = email
		SessionData.IsLogged = true
		SessionData.Error = ""
		fmt.Println("Registered : ", SessionData)

		CreateUser(GetDB(), username, password, email)
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
