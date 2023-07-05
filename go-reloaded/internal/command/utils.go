package command

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	commands        = map[string]string{"(hex)": "", "(bin)": "", "(up)": "", "(low)": "", "(cap)": ""}
	advCommands     = map[string]string{"(up,": "", "(low,": "", "(cap,": ""}
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
	return append(s[:index], s[index+1:]...)
}

func GetNum(s string) (output int, err error) {
	//FIXME: 12d 1d2. Only acceptable input should be "...digit)"
	expr := regexp.MustCompile(`\d+`)
	match := expr.FindString(s)
	if match == "" {
		return 0, ErrInvalidInput
	}
	output, err = strconv.Atoi(match)
	//if len(match) == 0 {
	//	return 0, ErrInvalidInput
	//}
	//fmt.Println(match)
	//output, err = strconv.Atoi(strings.Join(match, ""))
	if err != nil {
		return 0, err
	}
	return
}

//Example buzuk 30 books (cap,
//1
