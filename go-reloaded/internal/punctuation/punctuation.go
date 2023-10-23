package punctuation

import (
	"go-reloaded.aidostt.net/internal/command"
)

func Check(words *[]string) error {
	var (
		exist, isQuote       bool
		arr                  []string
		DQuoteCnt, SQuoteCnt int
	)

	// Loop through the words.
	for i := 0; i < len(*words); i++ {
		// Check for existence of the word in the regular expression map.
		_, exist = RegExp[rune((*words)[i][0])]
		// If the word exists, handle it based on the conditions.
		if exist {
			if i <= 0 {
				continue
			}
			if (*words)[i-1][0] == '(' {
				if i-2 < 0 {
					return command.ErrInvalidInput
				}
				(*words)[i-2] += (*words)[i]
				*words = delAtInd(*words, i)
				i -= 2
			} else {
				(*words)[i-1] += (*words)[i]
				*words = delAtInd(*words, i)
				i--
			}
		}
		// Check if the current word exists in the Quotes map.
		_, exist = Quotes[((*words)[i])]
		if exist {
			if isQuote {
				// If the previous word starts with an opening bracket, append the current word to it.
				if (*words)[i-1][0] == '(' {
					if !containsCommand((*words)[i-1]) {
						(*words)[i-1] += (*words)[i]
						// Remove the current word from the slice.
						*words = delAtInd(*words, i)
					} else {
						j := i - 1
						for j > 0 && containsCommand((*words)[j]) {
							j--
						}
						// Remove the current word from the slice.
						*words = delAtInd(*words, i)
					}
				} else {
					// If the previous word doesn't start with an opening bracket, append the current word to it.
					(*words)[i-1] += (*words)[i]
					// Remove the current word from the slice.
					*words = delAtInd(*words, i)
					// Update the loop counter.
					i--
					if i == 0 {
						continue
					}
				}
				isQuote = false
			} else {
				if i+1 >= len(*words) {
					return command.ErrInvalidInput
				}
				if (*words)[i+1][0] == '(' {
					if i+2 >= len(*words) {
						return command.ErrInvalidInput
					}
					(*words)[i+2] = (*words)[i] + (*words)[i+2]
					*words = delAtInd(*words, i)
				} else {
					(*words)[i+1] = (*words)[i] + (*words)[i+1]
					*words = delAtInd(*words, i)
				}
				isQuote = true
			}
		}
		//Call delimitWord function and check its result.
		if arr = delimitWord((*words)[i], &isQuote); len(arr) > 1 {
			// Make room for the new elements by extending the slice.
			*words = append(*words, make([]string, len(arr)-1)...)

			// Copy the elements from the end to the new position.
			copy((*words)[i+len(arr):], (*words)[i+1:])

			// Copy the elements from arr into words at the appropriate position.
			copy((*words)[i:i+len(arr)], arr)
			i -= 1
		}
	}
	for i := 0; i < len(*words); i++ {
		_, exist = Quotes[string((*words)[i][0])]
		if exist {
			for j := 0; j < len((*words)[i]); j++ {
				if (*words)[i] == "'" {
					SQuoteCnt++
				}
				if (*words)[i] == "\"" {
					DQuoteCnt++
				}
			}
		}
	}
	// Check if the counts of single and double Quotes are even.
	if SQuoteCnt%2 != 0 || DQuoteCnt%2 != 0 {
		return command.ErrInvalidInput
	}
	return nil
}
