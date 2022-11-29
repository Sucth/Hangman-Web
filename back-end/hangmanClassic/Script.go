package backend

type HangmanReturn struct {
	Lifes int
	LWord string
	Win   bool
}

func HangScript(life int, wordtofind string, wordanswer string, letter string) HangmanReturn {
	var result HangmanReturn
	var lwordtofind = []rune(wordtofind)
	var win bool
	var found = false
	if len(letter) != 0 {
		for i := range lwordtofind {
			if letter[0] == wordanswer[i] {
				found = true
				lwordtofind[i] = rune(wordanswer[i])
			}
		}
		if !found {
			life--
			result.Lifes = life
		} else {
			result.Lifes = life
		}
		win = true
		for _, i := range lwordtofind {
			if i == '_' {
				win = false
			}
		}
		result.Win = win
		result.LWord = string(lwordtofind)
		return result
	} else {
		result.Lifes = life
		result.Win = false
		result.LWord = string(lwordtofind)
		return result
	}
}
