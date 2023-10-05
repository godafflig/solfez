package utils

import (
	"html/template"
	"net/http"
)

// setting up session
func Login(email string, password string, w http.ResponseWriter, r *http.Request) {

	if !emailExists(GetDB(), email) {
		SessionData.Error = "Email incorrect."
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if !userExists(GetDB(), email, password) {
		SessionData.Error = "Mot de passe incorrect."
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		SessionData.Id = getId(GetDB(), email)
		SessionData.Username = getUsername(GetDB(), email)
		SessionData.Email = email
		SessionData.IsLogged = true
		SessionData.Score = getScore(GetDB(), email)
		SessionData.Error = ""
		SessionData.ProfilePic = GetProfilePicFromDb()
		template.Must(template.ParseFiles("static/Accueil.html")).Execute(w, SessionData)
	}
}

// creating a new session
func Register(username string, email string, password string, passwordCheck string, w http.ResponseWriter, r *http.Request) {
	if password != passwordCheck {
		SessionData.Error = "Les mots de passe ne correspondent pas."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// } else if !isStrongPassword(password) {
		// 	SessionData.Error = "Password must be at least 8 characters long and contain at least 1 digit, 1 symbol and 1 uppercase letter."
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if usernameExists(GetDB(), username) {
		SessionData.Error = "Le nom d'utilisateur existe déjà."
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if emailExists(GetDB(), email) {
		SessionData.Error = "L'email existe déjà."
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		CreateUser(GetDB(), username, password, email)
		SessionData.Id = getId(GetDB(), email)
		CreateScore(GetDB(), username, SessionData.Id)
		SessionData.Username = username
		SessionData.Email = email
		SessionData.IsLogged = true
		SessionData.Score = 0
		SessionData.Error = ""
		SessionData.ProfilePic = GetProfilePicFromDb()
		template.Must(template.ParseFiles("static/Accueil.html")).Execute(w, SessionData)
	}
}

// ending the current session and redirecting to home page
func Logout(w http.ResponseWriter, r *http.Request) {

	// change is_logged to 0 in database
	query := `
	UPDATE users SET is_logged = ? WHERE email = ?`
	_, err := GetDB().Exec(query, "0", SessionData.Email)
	if err != nil {
		panic(err)
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

// checking is the password is string enough
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
