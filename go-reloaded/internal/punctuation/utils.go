package punctuation

import (
	"strings"
)

var regExp = map[rune]rune{
	'!': ' ', '?': ' ', '.': ' ', ',': ' ',
	':': ' ', ';': ' ',
}

func delimitWord(word string) (out []string) {
	var (
		exist, sign bool
	)
	prevInd := 0
	rr := []rune(strings.TrimSpace(word))
	for i := 0; i < len(rr); i++ {
		_, exist = regExp[rr[i]]
		if exist {
			sign = true
		}
		if sign && !exist {
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
			//addSpaceAtInd(&rr, i)                    // Modify the rune slice directly.
			sign = false
			prevInd = i
		}
	}
	return append(out, string(rr[prevInd:]))
}

func addSpaceAtInd(rr *[]rune, i int) {
	*rr = append(*rr, 0) // Add a new rune to make space.
	copy((*rr)[i+1:], (*rr)[i:])
	(*rr)[i] = ' ' // Add a space rune.
}

func delAtInd(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
