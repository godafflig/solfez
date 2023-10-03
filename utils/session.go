package utils

type Session struct {
	Id       int
	Username string
	Email    string
	IsLogged bool
	Score    int
	Error    string
}

var SessionData Session
