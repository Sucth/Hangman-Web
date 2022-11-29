package backend

func HideWord(word string) string {
	var result = []rune(word)
	var nbtohide = len(word)/2 + 1

	for i := 0; i <= nbtohide; i++ {
		if result[i] == '_' {
			i--
		} else {
			result[i] = '_'
		}
	}

	return string(result)
}
