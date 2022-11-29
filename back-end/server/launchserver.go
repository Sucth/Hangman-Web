package backend

import (
	backend "HangmanWeb/back-end/hangmanClassic"
	"fmt"
	"html/template"
	"net/http"
)

func Launchserver() {
	http.Handle("/stylesheeets/", http.StripPrefix("/stylesheeets/", http.FileServer(http.Dir("Front-end/stylesheets"))))
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

type HangmanData struct {
	Lifes int
	LWord string
	Win   bool
}

var Liifes = 10
var Word = backend.ChooseWord()
var Lword = backend.HideWord(Word)
var Win = false

func EasyMode(w http.ResponseWriter, r *http.Request) {
	var Letter = r.FormValue("letter")
	result := backend.HangScript(Liifes, Lword, Word, Letter)
	Win = result.Win

	if Win == false {

		data := HangmanData{
			Lifes: result.Lifes,
			LWord: result.LWord,
		}
		Lword = result.LWord
		Liifes = result.Lifes

		var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
		err := templates.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
		err := templates.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
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
