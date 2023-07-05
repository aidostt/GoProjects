package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	commands        = map[string]string{"(hex)": "", "(bin)": "", "(up)": "", "(low)": "", "(cap)": ""}
	ErrCommNotFound = errors.New("no such command")
)

func ModificationText(file *os.File) ([]byte, error) {
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
				words[i-1] = strings.ToUpper(words[i-1])
			case "(low)":
				words[i-1] = strings.ToLower(words[i-1])
			case "(cap)":
				words[i-1] = strings.Title(words[i-1])
			default:
				return nil, ErrCommNotFound
			}
			words = DelAtInd(words, i)
		}
	}
	output := strings.Join(words, " ")
	return []byte(output), nil
}

func hex(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(n))
}

func bin(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(n))
}
