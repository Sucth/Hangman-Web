package backend

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func ChooseWord() string {
	file, err := ioutil.ReadFile("back-end/data/words/words.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Split(string(file), "\n")
	w := rand.Intn(len(words) - 1)
	word := words[w]
	return word
}
