package backend

func Wordtest(w string, r string) bool {
	word := []rune(w)
	wordtotest := []rune(r)
	if len(wordtotest) == len(word) {
		for i := 0; i <= len(word)-1; i++ {
			if word[i] == wordtotest[i] {
			} else {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
