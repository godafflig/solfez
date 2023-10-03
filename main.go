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
	}

}
