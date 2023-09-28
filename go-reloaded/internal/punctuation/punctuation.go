package punctuation

import (
	"go-reloaded.aidostt.net/internal/command"
)

func Check(words *[]string) error {
	var (
		exist      bool
		arr        []string
		DQuouteCnt int
		SQuouteCnt int
		s          string
	)
	for i := 0; i < len(*words); i++ {
		if (*words)[i] == "\"" {
			DQuouteCnt++
		}
		if (*words)[i] == "'" {
			SQuouteCnt++
		}

		if (*words)[i][0] == '(' {
			_, exist = regExp[rune((*words)[i][len((*words)[i])-1])]
			if exist {
				if i <= 0 {
					return command.ErrInvalidInput
				}
				for exist {
					s += string(rune((*words)[i][len((*words)[i])-1]))
					(*words)[i] = (*words)[i][:len((*words)[i])-1]
					_, exist = regExp[rune((*words)[i][len((*words)[i])-1])]
				}
				if (*words)[i-1][0] == '(' {
					for i >= 0 && (*words)[i-1][0] == '(' {
						i--
					}
					if i <= 0 {
						return command.ErrInvalidInput
					}
					(*words)[i-1] += s
					s = ""
				} else {
					(*words)[i-1] += s
					s = ""
				}
			}
		}
		_, exist = regExp[rune((*words)[i][0])]
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
		if arr = delimitWord((*words)[i], &SQuouteCnt, &DQuouteCnt); len(arr) > 1 {
			// Make room for the new elements by extending the slice.
			*words = append(*words, make([]string, len(arr)-1)...)

			// Copy the elements from the end to the new position.
			copy((*words)[i+len(arr):], (*words)[i+1:])

			// Copy the elements from arr into words at the appropriate position.
			copy((*words)[i:i+len(arr)], arr)
			i += len(arr) - 1
		}
	}
	if SQuouteCnt%2 != 0 || DQuouteCnt%2 != 0 {
		return command.ErrInvalidInput
	}
	return nil
}
