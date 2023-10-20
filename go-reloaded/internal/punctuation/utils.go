package punctuation

import (
	"strings"
)

// regExp defines a map of punctuation runes to be replaced with spaces.
var regExp = map[rune]rune{
	'!': ' ', '?': ' ', '.': ' ', ',': ' ',
	':': ' ', ';': ' ',
}

var quotes = map[string]rune{
	//"(": ' ', ")": ' ',
	"\"": ' ', "'": ' ',
}

// delimitWord function splits a word based on punctuation rules.
// It replaces punctuation with spaces and handles single and double quotes.
func delimitWord(word string, isQuote *bool) (out []string) {
	var (
		exist, sign, bracket bool
		sQuoteCnt, dQuoteCnt int
	)
	prevInd := 0
	rr := []rune(strings.TrimSpace(word))

	for i := 0; i < len(rr); i++ {
		_, exist = quotes[string(rr[i])]
		if exist {
			if rr[i] == '\'' {
				sQuoteCnt++
			} else {
				dQuoteCnt++
			}
		}
		_, exist = regExp[rr[i]]
		if exist {
			sign = true
		}

		// Check for closing bracket.
		if rr[i] == ')' {
			if i != len(rr)-1 {
				bracket = true
				continue
			}
		}

		// Handle bracketed words.
		if bracket {
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
			bracket = false
			prevInd = i
		}

		// Handle words with punctuation.
		if sign && !exist {
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
			sign = false
			prevInd = i
		}
	}

	// Append the last part of the word.
	out = append(out, string(rr[prevInd:]))

	// Handle special cases based on punctuation.
	if len(out) > 1 {
		for i := 0; i < len(out); i++ {
			_, exist = regExp[rune(out[i][0])]
			if exist {
				if i == 0 {
					continue
				}
				if out[i-1][len(out[i-1])-1] == ')' {
					j := i - 1
					for j > 0 && out[j][len(out[j])-1] == ')' {
						j--
					}
					s := out[j]
					out[j] = out[i]
					out = delAtInd(out, i)
					addWordAtInd(&out, j+1, s)
				}
			}
		}
	}
	if sQuoteCnt%2 != 0 || dQuoteCnt%2 != 0 {
		*isQuote = true
	}
	return out
}

// addWordAtInd inserts a word at a specific index in the slice.
func addWordAtInd(s *[]string, i int, word string) {
	*s = append(*s, "") // Add a new element to make space.
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = word
}

// delAtInd deletes an element at a specific index in the slice.
func delAtInd(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
