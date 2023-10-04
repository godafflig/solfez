package utils

import (
	"math/rand"
	"time"
)

// struct Game struct {
// 	Questions []string
// 	CorrectAnswer string
// }
//var gameData Game

var pianoKeys = []string{"do", "do#/réb", "ré", "ré#/mib", "mi", "fa", "fa#/solb", "sol", "sol#/lab", "la", "la#/sib", "si"}

func QuestionQCM() {
	var answers []string
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

	for i := 0; i < len(randomIndex); i++ {
		answers = append(answers, pianoKeys[randomIndex[i]])
	}

	// Choices
	//fmt.Println(answers)

	// Correct answer
	//correctAnswer := rand.Intn(3)
	//fmt.Println(answers[correctAnswer])

}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
