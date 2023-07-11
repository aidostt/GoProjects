package punctuation

func Check(words *[]string) error {
	for i := range *words {
		delimitWord(&(*words)[i])
		//if _, exist := regExp[rune(word[0])]; exist {
		//
		//}
		//switch case: identify when call right function
	}
	return nil
}

//example string:
//case 1: buzuk is example ...  buzuk is example ?! --> buzuk is example... buzuk is example?!
//case 2: " buzuk is example
//case 3: buzuk is example ...Is this real? --> buzuk is example... Is this real?
//case 4: buzuk is example;;  or .. or ,, or :: ERROR punctuation after ... ERROR
//case 5: If 1 " or ' ERROR
