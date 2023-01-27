package backend

import "math/rand"

func HideWord(word string) string {
	var result = []rune(word)
	var nbtohide = len(word)/2 + 1

	for i := 0; i <= nbtohide; i++ {
		n := rand.Intn(len(result) - 1)
		if result[n] == '_' {
			i--
		} else {
			result[n] = '_'
		}
	}

	return string(result)
}
