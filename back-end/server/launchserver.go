package backend

import (
	database "HangmanWeb/back-end/database"
	backend "HangmanWeb/back-end/hangmanClassic"
	"fmt"
	"html/template"
	"net/http"
)

func Launchserver() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("Front-end/stylesheets"))))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/difficultyPage", difficultyPage)
	http.HandleFunc("/Game", Game)
	http.HandleFunc("/Win", WinPage)
	http.HandleFunc("/Loose", Loose)
	http.HandleFunc("/leaderBoard", leaderBoard)
	fmt.Println("Server lancé sur le port 8080 au lien suivant : \n http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

var Pseudo string
var Liifes = 10
var Word string
var Lword string
var Win = false
var Lettersused []rune
var Message string
var GameStarted = false
var diff string

type HangmanData struct {
	Lifes      int
	LWord      string
	Msg        string
	Lettersusd string
	Difficulty string
}

func leaderBoard(w http.ResponseWriter, r *http.Request) {

	scores := database.SortScore(database.CollectScores())
	var templates = template.Must(template.ParseFiles("Front-end/templates/leaderboard.gohtml"))
	err := templates.Execute(w, scores)
	if err != nil {
		fmt.Println(err)
	}
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Front-end/templates/homePage.gohtml"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func difficultyPage(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("pseudo") != "" {
		Pseudo = r.PostFormValue("pseudo")
	}
	var templates = template.Must(template.ParseFiles("Front-end/templates/pageDifficulty.gohtml"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	if !GameStarted {
		if r.PostFormValue("difficulty") == "1" {
			Word = backend.ChooseWord("back-end/data/words/easyWords.txt")
			Lword = backend.HideWord(Word)
			diff = "1"
		} else if r.PostFormValue("difficulty") == "2" {
			Word = backend.ChooseWord("back-end/data/words/mediumWords.txt")
			Lword = backend.HideWord(Word)
			diff = "2"
		} else if r.PostFormValue("difficulty") == "3" {
			Word = backend.ChooseWord("back-end/data/words/hardWords.txt")
			Lword = backend.HideWord(Word)
			diff = "3"
		}
	}
	GameStarted = true
	var Letter = r.FormValue("letter")
	result := backend.Script(Liifes, Lword, Word, Letter, Lettersused)
	Win = result.Win
	Lword = result.LWord
	Liifes = result.Lifes
	Lettersused = result.Letters_used
	Message = result.Message

	if Win == false && Liifes > 0 {

		data := HangmanData{
			Lifes:      Liifes,
			LWord:      Lword,
			Lettersusd: string(Lettersused),
			Msg:        Message,
			Difficulty: diff,
		}

		var templates = template.Must(template.ParseFiles("Front-end/templates/pageGame.gohtml"))
		err := templates.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}

	} else if Win == true {
		http.Redirect(w, r, "/Win", http.StatusFound)

	} else if Liifes == 0 {
		http.Redirect(w, r, "/Loose", http.StatusFound)
	}

}

func WinPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "post" {

		http.Redirect(w, r, "/", http.StatusFound)
	}

	database.Addscore(Pseudo, diff)

	//resetvariables
	var resetpseudo string
	Pseudo = resetpseudo
	var resetltruseds []rune
	var resetdiff string
	Liifes = 10
	Win = false
	Lettersused = resetltruseds
	GameStarted = false
	diff = resetdiff
	//
	var templates = template.Must(template.ParseFiles("Front-end/templates/Win.gohtml"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Loose(w http.ResponseWriter, r *http.Request) {
	if r.Method == "post" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	//resetvariable
	var resetpseudo string
	Pseudo = resetpseudo
	var resetltruseds []rune
	var resetdiff string
	Liifes = 10
	Win = false
	Lettersused = resetltruseds
	GameStarted = false
	diff = resetdiff
	//
	var templates = template.Must(template.ParseFiles("Front-end/templates/Lose.gohtml"))
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
