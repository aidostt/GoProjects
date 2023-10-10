package punctuation

import (
	"strings"
)

var regExp = map[rune]rune{
	'!': ' ', '?': ' ', '.': ' ', ',': ' ',
	':': ' ', ';': ' ', '"': ' ', '\'': ' ',
}

func delimitWord(word string, SCnt *int, DCnt *int) (out []string) {
	var (
		exist, sign, bracket bool
	)
	prevInd := 0
	rr := []rune(strings.TrimSpace(word))
	for i := 0; i < len(rr); i++ {
		if rr[i] == '\'' {
			*SCnt++
		}
		if rr[i] == '"' {
			*DCnt++
		}
		_, exist = regExp[rr[i]]
		if exist {
			sign = true
		}
		if rr[i] == ')' {
			if i != len(rr)-1 {
				bracket = true
				continue
			}
		}
		if bracket {
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
			bracket = false
			prevInd = i
		}
		if sign && !exist {
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
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
