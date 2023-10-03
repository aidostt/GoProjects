package command

import (
	"errors"
	"strconv"
)

var (
	commands        = map[string]bool{"(hex)": false, "(bin)": false, "(up)": false, "(low)": false, "(cap)": false}
	advCommands     = map[string]bool{"(up,": false, "(low,": false, "(cap,": false}
	ErrCommNotFound = errors.New("no such command")
	ErrInvalidInput = errors.New("invalid input")
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
	return append((s)[:index], (s)[index+1:]...)
}

func number(s string) (output int, err error) {
	//Since in this function we expect to get the
	//string in format of "digit)", we need to get
	//rid of last bracket.
	s = s[:len(s)-1]
	output, err = strconv.Atoi(s)
	if err != nil {
		return 0, ErrInvalidInput
	}
	return
}

//TODO: handle multiple commands in a row something (up) (low) (cap)