package utils

import "fmt"

func UpdateStatistics(result string) {
	SessionData.Statistics.TotalGamesPlayed += 1
	UpdateTotalGames()
	if result == "win" {
		SessionData.Statistics.TotalGamesWon += 1
		UpdateWins()
	} else if result == "lose" {
		SessionData.Statistics.TotalGamesLost += 1
		UpdateLoses()
	}
	SessionData.HighestScore = GetScoreFromScoresTable()
}

func UpdateTotalGames() {
	db := GetDB()
	query := `
	UPDATE users SET total_games = ? WHERE user_id = ?`
	_, err := db.Exec(query, SessionData.Statistics.TotalGamesPlayed, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateWins() {
	db := GetDB()
	query := `
	UPDATE users SET wins = ? WHERE user_id = ?`
	_, err := db.Exec(query, SessionData.Statistics.TotalGamesWon, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateLoses() {
	db := GetDB()
	query := `
	UPDATE users SET loses = ? WHERE user_id = ?`
	_, err := db.Exec(query, SessionData.Statistics.TotalGamesLost, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
}
