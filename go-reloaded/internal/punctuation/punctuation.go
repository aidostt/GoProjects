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
	)
	//TODO: add command that will isolate all commands,
	//TODO: if they have any punctuation before and after
	for i := 0; i < len(*words); i++ {
		if (*words)[i] == "\"" {
			DQuouteCnt++
		}
		if (*words)[i] == "'" {
			SQuouteCnt++
		}
		_, exist = regExp[rune((*words)[i][0])]
		if exist {
			if i <= 0 {
				continue
			}
			//TODO: compare not first element, but the whole command
			if (*words)[i-1][0] == '(' {
				//a! (cap) !a
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
		//TODO: receive an array and put it into original array

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

//example string:
//case 1: buzuk is example ...  buzuk is example ?! --> buzuk is example... buzuk is example?!
//case 2: " buzuk is example
//case 3: buzuk is example ...Is this real? --> buzuk is example... Is this real?
//case 4: buzuk is example;;  or .. or ,, or :: ERROR punctuation after ... ERROR
//case 5: If 1 " or ' ERROR
