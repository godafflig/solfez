package utils

type Session struct {
	Id            int
	Username      string
	Email         string
	IsLogged      bool
	Score         int
	PreviousScore int
	HighestScore  int
	Error         string
	ProfilePic    string
	GameData      Game
	Statistics    Stats
}

type Game struct {
	Questions             []string
	CorrectAnswer         string
	CurrentLevel          int
	LifeLeft              int
	PreviousCorrectAnswer string
	CorrectNote           string
	Notes                 []string
}

type Stats struct {
	TotalGamesPlayed    int
	TotalGamesWon       int
	TotalGamesLost      int
	AccountCreatedSince string
	Time                string
	//TimeSpendPlaying    chan time.Duration
}

var SessionData Session

var GameData Game

var Statistics Stats
