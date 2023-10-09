package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

// upload the profile picture to the server & to the database
func HandleUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.Split(header.Filename, ".")
	header.Filename = uuid.New().String() + "." + ext[len(ext)-1]

	if SessionData.ProfilePic != "/static/assets/uploads/profil_placeholder.jpg" {
		os.Remove("." + SessionData.ProfilePic)
	}

	out, err := os.Create("./static/assets/uploads/" + header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SessionData.ProfilePic = "/static/assets/uploads/" + header.Filename

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = GetDB().Exec("UPDATE users SET profile_picture = ? WHERE user_id = ?", SessionData.ProfilePic, SessionData.Id)

	if err != nil {
		fmt.Errorf("failed to add profile pic artist: %v", err)
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

// get the profile picture from the database and returns it
func GetProfilePicFromDb() string {
	rows, err := GetDB().Query("SELECT profile_picture FROM users WHERE user_id = ?", SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var profilePic string
		if err := rows.Scan(&profilePic); err != nil {
			fmt.Println(err)
		}
		SessionData.ProfilePic = profilePic
	}
	return SessionData.ProfilePic
}

// change the username in the database
func ChangeUsername(oldUsername string, newUsername string, newUsernameConfirm string) {
	if UsernameExists(GetDB(), newUsername) {
		SessionData.Error = "Le nom d'utilisateur existe déjà."
	} else if newUsername != newUsernameConfirm {
		SessionData.Error = "Les deux nouveaux noms d'utilisateur ne correspondent pas."
	} else if UserExists(GetDB(), SessionData.Email, oldUsername) {
		UpdateUsername(GetDB(), SessionData.Email, newUsername)
		UpdateUsernameInScoresTable(GetDB(), SessionData.Id, newUsername)
		SessionData.Error = "Le nom d'utilisateur a été changé."
	} else {
		SessionData.Error = "Erreur. Vérifiez l'ancien nom d'utilisateur et essayez à nouveau."
	}
}

// change the password in the database
func ChangePassword(oldPassword string, newPassword string, newPasswordCheck string) {
	if newPassword == newPasswordCheck && UserExists(GetDB(), SessionData.Email, oldPassword) {
		UpdateUserPassword(GetDB(), SessionData.Email, newPassword)
		SessionData.Error = "Le mot de passe a été changé."
	} else {
		SessionData.Error = "Erreur. Vérifiez l'ancien mot de passe et essayez à nouveau."
	}
}

// clear the session data
func ClearDatas() {
	SessionData.Id = 0
	SessionData.Username = ""
	SessionData.Email = ""
	SessionData.IsLogged = false
	SessionData.Error = ""
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.Score = 0
	SessionData.Statistics.TotalGamesPlayed = 0
	SessionData.Statistics.TotalGamesWon = 0
	SessionData.Statistics.TotalGamesLost = 0
}
