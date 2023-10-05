package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var pianoKeys = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var pianoKeysDisplay = []string{"Do", "Do#", "Ré", "Ré#", "Mi", "Fa", "Fa#", "Sol", "Sol#", "La", "La#", "Si"}
var Octave = []string{"4", "5"}

func StartGame(w http.ResponseWriter, r *http.Request) {
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.CurrentLevel = 1
	InitializePathNotes()
	QuestionQCM(w, r)
}
func QuestionQCM(w http.ResponseWriter, r *http.Request) {

	var randomIndexNotes []int
	var randomIndexOctaves []int
	rand.Seed(time.Now().UnixNano())

	// Generate 3 random piano keys
	for j := 0; j < 3; j++ {

		for len(randomIndexNotes) <= 2 {
			n := rand.Intn(len(pianoKeysDisplay))
			if !contains(randomIndexNotes, n) {
				randomIndexNotes = append(randomIndexNotes, n)
			}
		}

		for len(randomIndexOctaves) <= 2 {
			o := rand.Intn(len(Octave))
			randomIndexOctaves = append(randomIndexOctaves, o)

		}
	}

	for i := 0; i < 3; i++ {
		SessionData.GameData.Questions = append(SessionData.GameData.Questions, pianoKeysDisplay[randomIndexNotes[i]]+Octave[randomIndexOctaves[i]]+"eme")
	}

	// Correct answer
	indexCorrectAnswer := rand.Intn(3)
	SessionData.GameData.CorrectNote = Octave[randomIndexOctaves[indexCorrectAnswer]] + pianoKeys[randomIndexNotes[indexCorrectAnswer]]
	fmt.Println(SessionData.GameData.CorrectNote)
	SessionData.GameData.CorrectAnswer = SessionData.GameData.Questions[indexCorrectAnswer]
	html := fmt.Sprintf(`

                <div id="Elnote" value="%s"></div>

        `, SessionData.GameData.CorrectNote)

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func CheckAnswer(answer string, w http.ResponseWriter, r *http.Request) bool {
	if answer == SessionData.GameData.CorrectAnswer {
		SessionData.Score += 1
		saveHighestScore(SessionData.Score)
		updateScore(GetDB(), SessionData.Email, SessionData.Score)
		SessionData.GameData.Questions = []string{}
		SessionData.GameData.CorrectAnswer = ""
		StartGame(w, r)
		return true
	} else {
		updateScore(GetDB(), SessionData.Email, SessionData.Score)
		SessionData.GameData.Questions = []string{}
		SessionData.GameData.CorrectAnswer = ""
		StartGame(w, r)
		http.Redirect(w, r, "/lost", http.StatusSeeOther)
		return false
	}
}

func InitializePathNotes() {
	for i := 0; i < len(Octave); i++ {
		for j := 0; j < len(pianoKeys); j++ {
			temp := Octave[i] + pianoKeys[j]
			SessionData.GameData.Notes = append(SessionData.GameData.Notes, temp)
		}
	}

}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
