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
		tmpl, _ := template.New("name").ParseFiles("static/index.html", "static/navbar.html")
		tmpl.ExecuteTemplate(w, "base", utils.SessionData)
		template.Must(template.ParseFiles("static/index.html")).Execute(w, utils.SessionData)
	case "/register-form":
		r.ParseForm()
		utils.Register(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("password-check"), w, r)
	case "/login":
		if r.Method == "GET" {
			template.Must(template.ParseFiles("static/login.html")).Execute(w, utils.SessionData)
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.Login(r.FormValue("email"), r.FormValue("password"), w, r)
		}
	case "/register":
		template.Must(template.ParseFiles("static/register.html")).Execute(w, utils.SessionData)
	case "/difficulte":
		utils.SessionData.Error = ""
		utils.SessionData.GameData.PreviousCorrectAnswer = ""
		ExecuteTemplate(w, r, "static/difficulte.html")
	case "/niveau-facile":
		if r.Method == "GET" {
			utils.StartGame(w, r, 1)
			ExecuteTemplate(w, r, "static/niveau-facile.html")
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			ExecuteTemplate(w, r, "static/niveau-facile.html")
		}
	case "/niveau-moyen":
		if r.Method == "GET" {
			utils.StartGame(w, r, 2)
			ExecuteTemplate(w, r, "static/niveau-moyen.html")
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			ExecuteTemplate(w, r, "static/niveau-moyen.html")
		}
	case "/niveau-difficile":
		if r.Method == "GET" {
			utils.StartGame(w, r, 3)
			ExecuteTemplate(w, r, "static/niveau-difficile.html")
		} else if r.Method == "POST" {
			r.ParseForm()
			utils.CheckAnswer(r.FormValue("answer"), w, r)
			ExecuteTemplate(w, r, "static/niveau-difficile.html")
		}
	case "/profile":
		utils.SessionData.Error = ""
		utils.SessionData.Statistics.AccountCreatedSince = utils.ConvertDateOfCreation()
		utils.ConvertDateOfCreation()
		if r.Method == "GET" {
			ExecuteTemplate(w, r, "static/profile.html")
		} else if r.Method == "POST" {
			utils.HandleUpload(w, r)
		}
	case "/update-username":
		r.ParseForm()
		utils.ChangeUsername(r.FormValue("oldusername"), r.FormValue("newusername"), r.FormValue("newusernameconfirm"))
		ExecuteTemplate(w, r, "static/profile.html")
	case "/update-password":
		r.ParseForm()
		utils.ChangePassword(r.FormValue("oldpassword"), r.FormValue("newpassword"), r.FormValue("newpwdconfirm"))
		ExecuteTemplate(w, r, "static/profile.html")
	case "/lost":
		ExecuteTemplate(w, r, "static/lost.html")
	case "/classement":
		utils.SortClassement()
		ExecuteTemplate(w, r, "static/classement.html")
	case "/logout":
		utils.Logout(w, r)
	case "/delete-account":
		utils.DeleteUser(utils.GetDB(), utils.SessionData.Email, w, r)
		template.Must(template.ParseFiles("static/index.html")).Execute(w, utils.SessionData)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, pageHtml string) {
	if utils.SessionData.IsLogged {
		tmpl, _ := template.New("name").ParseFiles(pageHtml, "static/navbar.html")
		tmpl.ExecuteTemplate(w, "base", utils.SessionData)
		template.Must(template.ParseFiles(pageHtml)).Execute(w, utils.SessionData)
	} else {
		utils.SessionData.Error = "Vous devez vous connecter pour accéder au jeu."
		template.Must(template.ParseFiles("static/login.html")).Execute(w, utils.SessionData)
	}
}
