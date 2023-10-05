package utils

import (
	"fmt"
	"sort"
	"strconv"
)

type Scoreboard struct {
	UserId     int
	Username   string
	Score      int
	Rank       int
	Classement []Scoreboard
}

var ScoreboardData Scoreboard

//var Classement []Scoreboard

func saveHighestScore(newScore int) {
	db := GetDB()

	// get score from score bdd
	query := `
	SELECT score FROM scores WHERE user_id = ?`
	rows, err := db.Query(query, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var oldScoreString string
	for rows.Next() {
		err := rows.Scan(&oldScoreString)
		if err != nil {
			fmt.Println(err)
		}
	}

	oldScore, err := strconv.Atoi(oldScoreString)
	if err != nil {
		fmt.Println(err)
	}

	if newScore > oldScore {
		query := `
		UPDATE scores SET score = ? WHERE user_id = ?`
		_, err := db.Exec(query, newScore, SessionData.Id)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SortClassement() {

	db := GetDB()

	// get all scores, usernames & id
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

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})
	ScoreboardData.Classement = scores

	// add the rank to the struct
	rank := 1
	for i := 0; i < len(ScoreboardData.Classement); i++ {
		ScoreboardData.Classement[i].Rank = rank
		rank++
	}
}
