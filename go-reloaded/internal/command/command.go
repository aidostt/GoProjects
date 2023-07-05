package command

import (
	"errors"
	"os"
	"strings"
)

func Check(file *os.File) ([]byte, error) {
	bytes := make([]byte, 2056)
	n, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(bytes[:n]), " ")

	for i, v := range words {
		if _, exist := commands[v]; exist {
			switch v {
			case "(hex)":
				words[i-1] = hex(words[i-1])
			case "(bin)":
				words[i-1] = bin(words[i-1])
			case "(up)":
				//FIXME: adapt this case for new func
				words[i-1] = strings.ToUpper(words[i-1])
			case "(low)":
				words[i-1] = strings.ToLower(words[i-1])
			case "(cap)":
				words[i-1] = strings.Title(words[i-1])
			default:
				return nil, ErrCommNotFound
			}
			words = delAtInd(words, i)
		}
	}
	num := 0
	for i, v := range words {
		if _, exist := advCommands[v]; exist {
			if words[i+1] == "" {
				return nil, ErrInvalidInput
			}
			num, err = GetNum(words[i+1])
			if err != nil {
				switch {
				case errors.Is(err, ErrInvalidInput):
					return nil, ErrInvalidInput
				default:
					return nil, ErrCommNotFound
				}
			}
			//if i+1 <= len(words) {
			//	num, err = GetNum(words[i+1])
			//	if err != nil {
			//		return nil, ErrCommNotFound
			//	}
			//} else {
			//	return nil, err
			//}
			switch v {
			case "(up,":
				for j := i; j >= i-num; j-- {
					//FIXME: fix the index out of range error
					words[j] = strings.ToUpper(words[j])
				}
			case "(low,":
				for j := i; j >= i-num; j-- {
					words[j] = strings.ToUpper(words[j])
				}
			case "(cap,":
				for j := i; j >= i-num; j-- {
					words[j] = strings.ToUpper(words[j])
				}
			}
			words = delAtInd(words, i)
			words = delAtInd(words, i)
		}
	}
	output := strings.Join(words, " ")
	return []byte(output), nil
}
