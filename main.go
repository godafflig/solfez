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
		//ServeMP3(w, r)
		//http.HandleFunc("/static/assets/C6.mp3", ServeMP3)

	case "/login":
		if r.Method == "GET" {
			template.Must(template.ParseFiles("static/login.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.Login(r.FormValue("email"), r.FormValue("password"), w, r)
		}
	case "/register":
		if r.Method == "GET" {
			template.Must(template.ParseFiles("static/register.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.Register(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("password-check"), w, r)
		}
	case "/play":
		if r.Method == "GET" {
			utils.StartGame(w, r)
			template.Must(template.ParseFiles("static/play.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			template.Must(template.ParseFiles("static/play.html")).Execute(w, utils.SessionData)
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
	}
}

func ServeMP3(w http.ResponseWriter, r *http.Request) {
	mp3FilePath := "./static/assets/C6.mp3"

	mp3File, err := os.Open(mp3FilePath)
	if err != nil {
		http.Error(w, "Fichier MP3 non trouvé", http.StatusNotFound)
		return
	}
	defer mp3File.Close()

	fileInfo, err := mp3File.Stat()
	if err != nil {
		http.Error(w, "Impossible d'obtenir la taille du fichier MP3", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")

	http.ServeContent(w, r, mp3FilePath, fileInfo.ModTime(), mp3File)
}
