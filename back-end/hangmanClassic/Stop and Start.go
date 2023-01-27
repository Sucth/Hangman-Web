package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Savetype struct {
	Word  string
	Lifes int
}

func Stop(lwords []rune, nblifes int) {
	word := string(lwords)
	save := Savetype{
		Word:  word,
		Lifes: nblifes,
	}

	file, err := json.Marshal(save)
	if err != nil {
		fmt.Print(err)
	}

	error := ioutil.WriteFile("save.txt", file, 0777)
	if error != nil {
		fmt.Print(error)
	}
}

func Start(filename string) (int, []rune) {
	var gamesave Savetype
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	error := json.Unmarshal(file, &gamesave)
	if error != nil {
		fmt.Print(error)
	}

	return gamesave.Lifes, []rune(gamesave.Word)
}
