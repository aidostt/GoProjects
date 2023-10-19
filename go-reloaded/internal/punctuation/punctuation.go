package punctuation

import (
	"go-reloaded.aidostt.net/internal/command"
)

func Check(words *[]string) error {
	var (
		exist                  bool
		arr                    []string
		DQuouteCnt, SQuouteCnt int
	)

	// Loop through the words.
	for i := 0; i < len(*words); i++ {
		// Count single and double quotes.
		if (*words)[i] == "\"" {
			DQuouteCnt++
		}
		if (*words)[i] == "'" {
			SQuouteCnt++
		}

		// Check for existence of the word in the regular expression map.
		_, exist = regExp[rune((*words)[i][0])]

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

		// Call delimitWord function and check its result.
		if arr = delimitWord((*words)[i], &SQuouteCnt, &DQuouteCnt); len(arr) > 1 {
			// Make room for the new elements by extending the slice.
			*words = append(*words, make([]string, len(arr)-1)...)

			// Copy the elements from the end to the new position.
			copy((*words)[i+len(arr):], (*words)[i+1:])

			// Copy the elements from arr into words at the appropriate position.
			copy((*words)[i:i+len(arr)], arr)
			i -= 1
		}
	}

	// Check if the counts of single and double quotes are even.
	if SQuouteCnt%2 != 0 || DQuouteCnt%2 != 0 {
		return command.ErrInvalidInput
	}

	return nil
}
