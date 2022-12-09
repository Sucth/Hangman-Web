package backend

type HangmanReturn struct {
	Lifes        int
	LWord        string
	Win          bool
	Message      string
	Letters_used string
}

func Script(life int, wordtofind string, wordanswer string, letter string, lettersused string) HangmanReturn {
	var Message string
	var Datas HangmanReturn
	var lifes = life
	var lwordtofind = []rune(wordtofind)
	var found = false
	var letter_used bool
	var letters_used = []rune(lettersused)

	if len(letter) > 1 {
		if wordanswer == letter {
			Datas.Win = true
		} else if !Wordtest(wordanswer, letter) {
			lifes -= 2
		}
	} else if len(letter) == 1 {
		letter_used = false
		if len(letters_used) == 0 {
			letters_used = append(letters_used, rune(letter[0]))
			letters_used = append(letters_used, ',')
		} else {
			for i := 0; i < len(letters_used); i++ {
				if letters_used[i] == rune(letter[0]) {
					letter_used = true
					break
				}
			}
		}
		if letter_used {
			Message = "You already used this letter"
		} else {
			letters_used = append(letters_used, rune(letter[0]))
			letters_used = append(letters_used, ',')
		}
		for i := 0; i < len(wordanswer); i++ {
			if rune(wordanswer[i]) == rune(letter[0]) {
				lwordtofind[i] = rune(letter[0])
				found = true
			}
		}
		if found {

		} else {
			Message = "Wrong letter (-1 life)"
			lifes--
		}
		if !letter_used {
			Datas.Win = true
			for _, n := range lwordtofind {
				if n == '_' {
					Datas.Win = false
				}
			}
		} else {
			if found {
				Message = "You already used this letter"
			} else {
				lifes--
				Message = "Wrong letter"
			}
		}
	}
	Datas.Lifes = lifes
	Datas.LWord = string(lwordtofind)
	Datas.Message = Message
	Datas.Letters_used = string(letters_used)
	return Datas
}
