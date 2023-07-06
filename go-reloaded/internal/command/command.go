package command

import (
	"strings"
)

func Check(words []string) ([]byte, error) {
	for i, v := range words {
		if _, exist := commands[v]; exist {
			switch v {
			case "(hex)":
				words[i-1] = hex(words[i-1])
			case "(bin)":
				words[i-1] = bin(words[i-1])
			case "(up)":
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
	var (
		num int
		err error
	)
	for i, v := range words {
		if _, exist := advCommands[v]; exist {
			if i+1 > len(words)-1 {
				return nil, ErrInvalidInput
			}
			num, err = number(words[i+1])
			if err != nil {
				return nil, ErrInvalidInput
			}
			switch v {
			case "(up,":
				for j := i; j >= i-num; j-- {
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
