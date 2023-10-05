package utils

import (
	"math/rand"
	"net/http"
	"time"
)

var pianoKeys = []string{"do", "do#/réb", "ré", "ré#/mib", "mi", "fa", "fa#/solb", "sol", "sol#/lab", "la", "la#/sib", "si"}

func StartGame(w http.ResponseWriter, r *http.Request) {
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.Questions = []string{}
    SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.CurrentLevel = 1
	QuestionQCM()
}
func QuestionQCM() {

	var randomIndex []int
	rand.Seed(time.Now().UnixNano())

	// Generate 3 random piano keys
	for j := 0; j < 3; j++ {

		for len(randomIndex) != 3 {
			n := rand.Intn(len(pianoKeys) - 1)
			if !contains(randomIndex, n) {
				randomIndex = append(randomIndex, n)
			}
		}
	}

	for i := 0; i < 3; i++ {
		SessionData.GameData.Questions = append(SessionData.GameData.Questions, pianoKeys[randomIndex[i]])
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

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
