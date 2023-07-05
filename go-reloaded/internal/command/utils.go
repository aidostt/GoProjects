package command

import (
	"errors"
	"strconv"
)

var (
	commands        = map[string]string{"(hex)": "", "(bin)": "", "(up)": "", "(low)": "", "(cap)": ""}
	ErrCommNotFound = errors.New("no such command")
)

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

func delAtInd(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
