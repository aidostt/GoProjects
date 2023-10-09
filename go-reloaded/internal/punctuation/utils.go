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
		exist, sign bool
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
		if sign && !exist {
			if rr[i] == ')' {
				continue
			}
			out = append(out, string(rr[prevInd:i])) // Append the delimited word.
			sign = false
			prevInd = i
		}

		//if rr[i] == '(' {
		//	//TODO:handle punctuation before and after command (advanced too)
		//	//TODO: something Very! (up) !(cap, 2)!!
		//	if i != 0 {
		//		fmt.Printf("rr --> %v\n", string(rr[:i]))
		//		//if before, cut from the beginning till the i and add it to out string
		//		out = append(out, string(rr[:i]))
		//	}
		//
		//	//_, exist = regExp[rune((*words)[i][len((*words)[i])-1])]
		//	//			if exist {
		//	//				if i <= 0 {
		//	//					return command.ErrInvalidInput
		//	//				}
		//	//				for exist {
		//	//					s += string(rune((*words)[i][len((*words)[i])-1]))
		//	//					(*words)[i] = (*words)[i][:len((*words)[i])-1]
		//	//					_, exist = regExp[rune((*words)[i][len((*words)[i])-1])]
		//	//				}
		//	//				if (*words)[i-1][0] == '(' {
		//	//					for i >= 0 && (*words)[i-1][0] == '(' {
		//	//						i--
		//	//					}
		//	//					if i <= 0 {
		//	//						return command.ErrInvalidInput
		//	//					}
		//	//					(*words)[i-1] += s
		//	//					s = ""
		//	//				} else {
		//	//					(*words)[i-1] += s
		//	//					s = ""
		//	//				}
		//	//			}
		//	//this peace of code can help in future
		//
		//}
		//if rr[i] == ')' {
		//	if i != len(rr)-1 {
		//		out = append(out, string(rr[i:]))
		//		//if after, cut from i till the end or next bracket and add it to out string
		//
		//	}
		//}
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
