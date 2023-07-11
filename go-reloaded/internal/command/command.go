package command

import (
	"strings"
)

func Check(words *[]string) error {
	var exist bool
	l := len(*words)
	for i := l - 1; i >= 0; i-- {
		if _, exist = commands[(*words)[i]]; exist {
			switch (*words)[i] {
			case "(hex)":
				(*words)[i-1] = hex((*words)[i-1])
			case "(bin)":
				(*words)[i-1] = bin((*words)[i-1])
			case "(up)":
				(*words)[i-1] = strings.ToUpper((*words)[i-1])
			case "(low)":
				(*words)[i-1] = strings.ToLower((*words)[i-1])
			case "(cap)":
				(*words)[i-1] = strings.Title((*words)[i-1])
			default:
				return ErrCommNotFound
			}
			*words = delAtInd(*words, i)
			i--
			l--
		}
	}
	var (
		num int
		err error
	)
	for i := l - 1; i >= 0; i-- {
		if _, exist := advCommands[(*words)[i]]; exist {
			if i+1 > len(*words)-1 {
				return ErrInvalidInput
			}
			num, err = number((*words)[i+1])
			if err != nil {
				return ErrInvalidInput
			}
			switch (*words)[i] {
			case "(up,":
				for j := i; j >= i-num; j-- {
					(*words)[j] = strings.ToUpper((*words)[j])
				}
			case "(low,":
				for j := i; j >= i-num; j-- {
					(*words)[j] = strings.ToUpper((*words)[j])
				}
			case "(cap,":
				for j := i; j >= i-num; j-- {
					(*words)[j] = strings.ToUpper((*words)[j])
				}
			}
			*words = delAtInd(*words, i)
			i--
			l--
			*words = delAtInd(*words, i)
			i--
			l--
		}
	}
	return nil
}
