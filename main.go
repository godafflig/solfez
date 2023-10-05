package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"utils/utils"

	"github.com/joho/godotenv"
)

func main() {
	// creating database if not exist
	utils.CreateUserTable(utils.GetDB())
	utils.CreateScoreTable(utils.GetDB())
	utils.SortClassement()

	// loading port & url from .env file
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("./.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	PORT := os.Getenv("PORT")
	URL := os.Getenv("URL")

	// static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// starting server
	http.HandleFunc("/", Routing)
	log.Println("Listening on " + URL + ":" + PORT)
	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Routing(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/":
		template.Must(template.ParseFiles("static/index.html")).Execute(w, utils.SessionData)
	case "/register":
		r.ParseForm()
		utils.Register(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("password-check"), w, r)
	case "/login":
		if r.Method == "GET" {
			template.Must(template.ParseFiles("static/login.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.Login(r.FormValue("email"), r.FormValue("password"), w, r)
		}
	case "/niveau-facile":
		if r.Method == "GET" {
			utils.StartGame(w, r, 1)
			template.Must(template.ParseFiles("static/niveau-facile.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			template.Must(template.ParseFiles("static/niveau-facile.html" )).Execute(w, utils.SessionData)
		}
	case "/niveau-moyen":
		if r.Method == "GET" {
			utils.StartGame(w, r, 2)
			template.Must(template.ParseFiles("static/niveau-moyen.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			template.Must(template.ParseFiles("static/niveau-moyen.html")).Execute(w, utils.SessionData)
		}
	case "/niveau-difficile":
		if r.Method == "GET" {
			utils.StartGame(w, r, 2)
			template.Must(template.ParseFiles("static/niveau-difficile.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			template.Must(template.ParseFiles("static/niveau-difficile.html")).Execute(w, utils.SessionData)
		}
	case "/profile":
		if r.Method == "GET" {
			template.Must(template.ParseFiles("static/profile.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			utils.HandleUpload(w, r)
		}
	case "/logout":
		utils.Logout(w, r)
	case "/lost":
		template.Must(template.ParseFiles("static/lost.html")).Execute(w, utils.SessionData)
	case "/classement":
		utils.SortClassement()
		template.Must(template.ParseFiles("static/classement.html")).Execute(w, utils.ScoreboardData)
	case "/accueil":
		template.Must(template.ParseFiles("static/Accueil.html")).Execute(w, utils.SessionData)
	case "/difficulte":
		template.Must(template.ParseFiles("static/difficulte.html")).Execute(w, utils.SessionData)
	case "/delete-account":
		utils.DeleteUser(utils.GetDB(), utils.SessionData.Email)
		template.Must(template.ParseFiles("static/index.html")).Execute(w, utils.SessionData)
	}
}
