package punctuation

import (
	"fmt"
	"go-reloaded.aidostt.net/internal/command"
)

func Check(words *[]string) error {
	var (
		exist bool
		arr   []string
	)
	//TODO: add command that will isolate all commands,
	//TODO: if they have any punctuation before and after
	for i := 0; i < len(*words); i++ {
		//fmt.Println(i)
		//fmt.Println(*words)
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
		if arr = delimitWord((*words)[i]); len(arr) > 1 {
			temp := make([]string, len(arr)+len(*words)-1)
			fmt.Printf("stage 1 ->%v\n", temp)

			temp = append((*words)[:i], arr...)
			fmt.Printf("stage 2 ->%v\n", temp)

			temp = append(temp, (*words)[i+1:]...)
			fmt.Printf("stage 3 ->%v\n", temp)
			fmt.Printf("words -> %v\n", (*words)[i+1:])
			*words = temp
			temp = nil
			i += len(arr) - 1
			//paste new arr into *words
			//increase the i position by len of the new array
		}
	}
	return nil
}

//example string:
//case 1: buzuk is example ...  buzuk is example ?! --> buzuk is example... buzuk is example?!
//case 2: " buzuk is example
//case 3: buzuk is example ...Is this real? --> buzuk is example... Is this real?
//case 4: buzuk is example;;  or .. or ,, or :: ERROR punctuation after ... ERROR
//case 5: If 1 " or ' ERROR
