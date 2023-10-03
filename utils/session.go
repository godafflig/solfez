package utils

type Session struct {
	Id       int
	IsLogged bool
	Error    string
}

var SessionData Session
