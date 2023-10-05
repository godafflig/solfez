package utils

import (
	"math/rand"
	"net/http"
	"time"
)

var pianoKeys = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var Octave = []string{"4", "5"}

func StartGame(w http.ResponseWriter, r *http.Request) {
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.Questions = []string{}
    SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.CurrentLevel = 1
	InitializePathNotes()
	QuestionQCM()
}
func QuestionQCM() {

	var randomIndexNotes []int
	var randomIndexOctaves []int
	rand.Seed(time.Now().UnixNano())

	// Generate 3 random piano keys
	for j := 0; j < 3; j++ {

		for len(randomIndexNotes) <= 2 {
			n := rand.Intn(len(pianoKeys))
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
		SessionData.GameData.Questions = append(SessionData.GameData.Questions, pianoKeys[randomIndexNotes[i]]+Octave[randomIndexOctaves[i]]+"eme")
	}

	// Correct answer
	indexCorrectAnswer := rand.Intn(3)
	SessionData.GameData.CorrectAnswer = SessionData.GameData.Questions[indexCorrectAnswer]
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
