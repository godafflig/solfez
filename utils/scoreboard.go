package utils

import (
	"fmt"
	"sort"
)

// compare & save the highest score in the database 'scores'
func saveHighestScore(newScore int) {
	db := GetDB()
	oldScore := GetScoreFromScoresTable()
	if newScore > oldScore {
		UpdateScoreInScoreTable(db, SessionData.Id, newScore)
	}
}

// sort the classement by score and add the rank
func SortClassement() {
	db := GetDB()

	// Clear the previous data
	ScoreboardData.Classement = []Scoreboard{}

	// define the IsLogged value
	ScoreboardData.IsLogged = true

	// get all scores, usernames, profile pics & id from 'scores' table
	query := `SELECT user_id, user_name, score FROM scores`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var scores []Scoreboard

	for rows.Next() {
		var userID, scoreValue int
		var username string
		if err := rows.Scan(&userID, &username, &scoreValue); err != nil {
			fmt.Println(err)
			continue
		}
		scores = append(scores, Scoreboard{UserId: userID, Username: username, Score: scoreValue})
	}

	// sort the scores by score
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	SessionData.ScoreboardData.Classement = scores

	// add the rank to the struct
	rank := 1
	for i := 0; i < len(SessionData.ScoreboardData.Classement); i++ {
		SessionData.ScoreboardData.Classement[i].Rank = rank
		rank++
	}
}
