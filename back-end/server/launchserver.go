package backend

import (
	backend "HangmanWeb/back-end/hangmanClassic"
	"fmt"
	"html/template"
	"net/http"
)

func Launchserver() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("Front-end/stylesheets"))))
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
	Lifes      int
	LWord      string
	Win        bool
	Lettersusd string
}

var Liifes = 10
var Word = backend.ChooseWord()
var Lword = backend.HideWord(Word)
var Win = false
var Lettersused string

func EasyMode(w http.ResponseWriter, r *http.Request) {
	var Letter = r.FormValue("letter")
	result := backend.Script(Liifes, Lword, Word, Letter, Lettersused)
	Win = result.Win
	Lword = result.LWord
	Liifes = result.Lifes

	if Win == false && Liifes > 0 {

		data := HangmanData{
			Lifes:      Liifes,
			LWord:      Lword,
			Lettersusd: Lettersused,
		}

		var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
		err := templates.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}

	} else if Win == true {
		data := HangmanData{Win: true, LWord: "You Won !"}
		var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
		err := templates.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		data := HangmanData{Lifes: 0, LWord: "You lost"}
		var templates = template.Must(template.ParseFiles("Front-end/templates/pageEasy.html"))
		err := templates.Execute(w, data)
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
