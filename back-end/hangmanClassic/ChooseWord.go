package backend

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func ChooseWord(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	list_word_str := strings.Split(string(file), "\r\n")

	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(len(list_word_str))
	result := list_word_str[v]
	return result
}
