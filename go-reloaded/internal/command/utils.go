package command

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

var quotes = map[string]rune{
	"(": ' ', ")": ' ',
	"\"": ' ', "'": ' ',
}

func containsCommand(word string) bool {
	for i := 0; i < len(word); i++ {
		for command := range Сommands {
			if i+len(command) <= len(word) && word[i:i+len(command)] == command {
				return true
			}
		}
	}
	return false
}

func hexToDec(hexStr string) string {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return hexStr // Return the original string if conversion fails.
	}
	return fmt.Sprintf("%d", dec)
}

// hex function converts a hexadecimal string to a decimal string.
func hex(input string) string {
	re := regexp.MustCompile(`[0-9A-Fa-f]+`) // Regular expression to find hexadecimal sequences.
	hexStrings := re.FindAllString(input, -1)
	fmt.Println(input)
	if hexStrings == nil {
		return input // No hex sequences found, return the original string.
	}

	// Replace each hexadecimal string with its decimal equivalent.
	for _, hexStr := range hexStrings {
		decStr := hexToDec(hexStr)
		input = strings.Replace(input, hexStr, decStr, 1)
	}

	return input
}

func binToDec(binStr string) string {
	dec, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr // Return the original string if conversion fails.
	}
	return fmt.Sprintf("%d", dec)
}

// bin function converts a binary string to a decimal string.
func bin(input string) string {
	re := regexp.MustCompile(`[01]+`) // Regular expression to find binary sequences.
	binStrings := re.FindAllString(input, -1)

	if binStrings == nil {
		return input // No binary sequences found, return the original string.
	}

	// Replace each binary string with its decimal equivalent.
	for _, binStr := range binStrings {
		decStr := binToDec(binStr)
		input = strings.Replace(input, binStr, decStr, 1)
	}

	return input
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
