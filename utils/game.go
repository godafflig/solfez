package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var pianoKeys = []string{"C", "CT", "D", "DT", "E", "F", "FT", "G", "GT", "A", "AT", "B"}
var pianoKeysDisplay = []string{"Do", "Do#", "Ré", "Ré#", "Mi", "Fa", "Fa#", "Sol", "Sol#", "La", "La#", "Si"}
var Octave = []string{"4", "5"}

// initiate the game depending on the level
func StartGame(w http.ResponseWriter, r *http.Request, level int) {
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.Error = ""

	switch level {
	case 1:
		SessionData.GameData.CurrentLevel = 1
		SessionData.GameData.LifeLeft = 3
	case 2:
		SessionData.GameData.CurrentLevel = 2
		SessionData.GameData.LifeLeft = 2
	case 3:
		SessionData.GameData.CurrentLevel = 3
		SessionData.GameData.LifeLeft = 1
	}
	QuestionQCM(w, r)
}

// continue playing the game
func PlayAgain(w http.ResponseWriter, r *http.Request, lifeleft int, level int) {
	SessionData.GameData.Questions = []string{}
	SessionData.GameData.CorrectAnswer = ""
	SessionData.GameData.CurrentLevel = level
	SessionData.GameData.LifeLeft = lifeleft
	QuestionQCM(w, r)
}

// creating 3 questions & one correct answer
func QuestionQCM(w http.ResponseWriter, r *http.Request) {

	var randomIndexNotes []int
	var randomIndexOctaves []int
	rand.Seed(time.Now().UnixNano())

	// Generate 3 random piano keys
	for j := 0; j < 3; j++ {

		for len(randomIndexNotes) != 3 {
			n := rand.Intn(len(pianoKeys) - 1)
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

	indexCorrectAnswer := rand.Intn(3)
	SessionData.GameData.CorrectNote = Octave[randomIndexOctaves[indexCorrectAnswer]] + pianoKeys[randomIndexNotes[indexCorrectAnswer]]
	SessionData.GameData.CorrectAnswer = SessionData.GameData.Questions[indexCorrectAnswer]
	html := fmt.Sprintf(`

	<div id="Elnote" value="%s"></div>

`, SessionData.GameData.CorrectNote)

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// checking if the answer is correct and updating the datas accordlingly
func CheckAnswer(answer string, w http.ResponseWriter, r *http.Request) bool {

	if answer == SessionData.GameData.CorrectAnswer {
		switch SessionData.GameData.CurrentLevel {
		case 1:
			SessionData.Score += 1
		case 2:
			SessionData.Score += 5
		case 3:
			SessionData.Score += 10
		}
		SessionData.Error = "Youpi tu l'as trouvé ! :)"
		UpdateScore(GetDB(), SessionData.Email, SessionData.Score)
		saveHighestScore(SessionData.Score)
		SortClassement()
		UpdateStatistics("win")
		SessionData.GameData.Questions = []string{}
		SessionData.GameData.PreviousCorrectAnswer = "La bonne réponse était : " + SessionData.GameData.CorrectAnswer
		SessionData.GameData.CorrectAnswer = ""
		PlayAgain(w, r, SessionData.GameData.LifeLeft, SessionData.GameData.CurrentLevel)
		return true
	} else {
		UpdateScore(GetDB(), SessionData.Email, SessionData.Score)
		saveHighestScore(SessionData.Score)
		SortClassement()
		SessionData.Error = "Oups... Essaie encore !"
		UpdateStatistics("lose")
		SessionData.GameData.Questions = []string{}
		SessionData.GameData.PreviousCorrectAnswer = "La bonne réponse était : " + SessionData.GameData.CorrectAnswer
		SessionData.GameData.CorrectAnswer = ""
		if SessionData.GameData.LifeLeft > 1 {
			SessionData.GameData.LifeLeft -= 1
			PlayAgain(w, r, SessionData.GameData.LifeLeft, SessionData.GameData.CurrentLevel)
		} else {
			SessionData.Score = 0
			UpdateScore(GetDB(), SessionData.Email, SessionData.Score)
			http.Redirect(w, r, "/lost", http.StatusSeeOther)
		}
		return false
	}
}

// check if a value is in an array
func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func InitializePathNotes() {
	for i := 0; i < len(Octave); i++ {
		for j := 0; j < len(pianoKeys); j++ {
			temp := Octave[i] + pianoKeys[j]
			SessionData.GameData.Notes = append(SessionData.GameData.Notes, temp)
		}
	}

}
