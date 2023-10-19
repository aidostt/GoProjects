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
	out = append(out, string(rr[prevInd:]))
	if len(out) == 1 {
		return nil
	} else {
		for i := 0; i < len(out); i++ {
			_, exist = regExp[rune(out[i][0])]
			if exist {
				if i == 0 {
					continue
				}
				if out[i-1][len(out[i-1])-1] == ')' {
					j := i - 1
					for out[j][len(out[j])-1] == ')' {
						if j <= 0 {
							break
						} else {
							j--
						}
					}
					s := out[j]
					out[j] = out[i]
					out = delAtInd(out, i)
					addWordAtInd(&out, j+1, s)
				}
			}
		}
	}
	return out
}

func addWordAtInd(s *[]string, i int, word string) {
	*s = append(*s, "") // Add a new rune to make space.
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = word // Add a space rune.
}

func delAtInd(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
