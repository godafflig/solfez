package utils

type Session struct {
	Id       int
	Username string
	Email    string
	IsLogged bool
	Error    string
}

var SessionData Session
