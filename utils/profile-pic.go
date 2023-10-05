package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

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
