package punctuation

var regExp = map[rune]rune{
	'!': ' ', '?': ' ', '.': ' ', ',': ' ',
	':': ' ', ';': ' ',
}

func delimitWord(word *string) {
	var exist, sign bool
	for i, el := range *word {
		_, exist = regExp[el]
		if exist {
			sign = true
		}
		if sign && !exist {
			addSpaceAtInd(word, i)
			sign = false
		}
	}
}

func addSpaceAtInd(s *string, i int) {
	*s += " "
	rr := []rune(*s)
	copy(rr[i+1:], rr[i:])
	rr[i] = ' '
	*s = string(rr)
}

func delAtInd(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
