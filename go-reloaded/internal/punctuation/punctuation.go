package punctuation

import (
	"go-reloaded.aidostt.net/internal/command"
)

func Check(words *[]string) error {
	var exist bool
	//TODO: add command that will isolate all commands,
	//TODO: if they have any punctuation before and after
	for i := range *words {
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

		delimitWord(&(*words)[i])
		//switch case: identify when call right function
	}
	return nil
}

//string!!!! buzuk

//example string:
//case 1: buzuk is example ...  buzuk is example ?! --> buzuk is example... buzuk is example?!
//case 2: " buzuk is example
//case 3: buzuk is example ...Is this real? --> buzuk is example... Is this real?
//case 4: buzuk is example;;  or .. or ,, or :: ERROR punctuation after ... ERROR
//case 5: If 1 " or ' ERROR
