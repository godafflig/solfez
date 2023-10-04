package utils

import "fmt"

type Scoreboard struct {
}

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

	//comparer le score dans la bdd user et le score dans l'autre bdd
	// enregister le highest


}

func classement() {
	// mettre dans un tableau qui sera dans une struct dans l'ordre
	scores := []int{5, 4, 6, 6, 55, 1}
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}
