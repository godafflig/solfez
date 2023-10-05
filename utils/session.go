package utils

type Session struct {
	Id         int
	Username   string
	Email      string
	IsLogged   bool
	Score      int
	Error      string
	ProfilePic string
	GameData   Game
}

type Game struct {
	Questions     []string
	CorrectAnswer string
	CurrentLevel  int
	Notes         []string
	CorrectNote   string
}

var SessionData Session

var GameData Game
