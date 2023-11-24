package internal

import (
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Colorize(alphabet map[rune]string, ltrToClrz, color string) (map[rune]string, error) {
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"gray":   "\033[37m",
		"white":  "\033[38m",
		"reset":  "\033[0m",
	}
	//TODO:bonus task: implement the rgb logic
	for _, el := range ltrToClrz {
		alphabet[el] = colors[strings.ToLower(color)] + alphabet[el] + colors["reset"]

	}
	return alphabet, nil
}
