package command

import (
	"errors"
	"strconv"
)

var (
	// Сommands represents a set of valid single-word commands.
	Сommands = map[string]bool{"(hex)": false, "(bin)": false, "(up)": false, "(low)": false, "(cap)": false}

	// advCommands represents a set of valid multi-word commands.
	advCommands = map[string]bool{"(up,": false, "(low,": false, "(cap,": false}

	// ErrCommNotFound is returned when a command is not found in the commands map.
	ErrCommNotFound = errors.New("no such command")

	// ErrInvalidInput is returned when the input is invalid or doesn't meet the expected format.
	ErrInvalidInput = errors.New("invalid input")
)

// hex function converts a hexadecimal string to a decimal string.
func hex(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(n))
}

// bin function converts a binary string to a decimal string.
func bin(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(n))
}

// delAtInd function deletes an element at a specific index in a string slice.
func delAtInd(s []string, index int) []string {
	return append((s)[:index], (s)[index+1:]...)
}

// number function extracts a number from a string formatted as "digit)" and returns it as an integer.
func number(s string) (output int, err error) {
	// Since this function expects the string in the format "digit)", the last bracket is removed.
	s = s[:len(s)-1]
	output, err = strconv.Atoi(s)
	if err != nil {
		return 0, ErrInvalidInput
	}
	return
}
