package backend

import (
	"fmt"
	"html/template"
	"net/http"
)

func Launchserver() {
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/difficultyPage", difficultyPage)
	http.HandleFunc("/EasyMode", EasyMode)
	http.HandleFunc("/MediumMode", MediumMode)
	http.HandleFunc("/HardMode", HardMode)
	fmt.Println("Server lancé sur le port 8080 au lien suivant : \n http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func Accueil(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/homePage.html"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func difficultyPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/pageDifficulty.html"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func EasyMode(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func MediumMode(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/pageMedium.html"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func HardMode(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/pageHard.html"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
