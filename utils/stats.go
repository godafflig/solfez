package utils

import (
	"fmt"
	"strconv"
)

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

func GetTotalGamesPlayed() int {
	db := GetDB()
	query := `
	SELECT total_games FROM users WHERE user_id = ?`
	rows, err := db.Query(query, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var totalGamesStr string
	for rows.Next() {
		err := rows.Scan(&totalGamesStr)
		if err != nil {
			fmt.Println(err)
		}
	}
	totalGames, err := strconv.Atoi(totalGamesStr)
	if err != nil {
		fmt.Println(err)
	}
	return totalGames
}

func GetTotalGamesWon() int {
	db := GetDB()
	query := `
	SELECT wins FROM users WHERE user_id = ?`
	rows, err := db.Query(query, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var wonStr string
	for rows.Next() {
		err := rows.Scan(&wonStr)
		if err != nil {
			fmt.Println(err)
		}
	}

	wonInt, err := strconv.Atoi(wonStr)
	if err != nil {
		fmt.Println(err)
	}
	return wonInt
}

func GetTotalGamesLost() int {
	db := GetDB()
	query := `
	SELECT loses FROM users WHERE user_id = ?`
	rows, err := db.Query(query, SessionData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var losesStr string
	for rows.Next() {
		err := rows.Scan(&losesStr)
		if err != nil {
			fmt.Println(err)
		}
	}

	losesInt, err := strconv.Atoi(losesStr)
	if err != nil {
		fmt.Println(err)
	}
	return losesInt
}
