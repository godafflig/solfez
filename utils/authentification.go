package utils

import (
	"html/template"
	"net/http"
)

func Login(email string, password string, w http.ResponseWriter, r *http.Request) {

	if CheckIfUserExist(GetDB(), email, password) {
		template.Must(template.ParseFiles("static/ok.html")).Execute(w, SessionData)
	} else {
		SessionData.Error = "Bad credentials."
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Register(username string, email string, password string, passwordCheck string, w http.ResponseWriter, r *http.Request) {

	if password == passwordCheck {
		CreateUser(GetDB(), username, password, email)
		template.Must(template.ParseFiles("static/ok.html")).Execute(w, SessionData)
	} else {
		SessionData.Error = "Bad credentials."
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}
