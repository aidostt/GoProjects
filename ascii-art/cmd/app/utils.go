package main

import (
	"errors"
)

func assignArgs(args []string) (input, lToColorize, desiredFont string, err error) {
	switch len(args) {
	case 1:
		input, desiredFont = args[0], STANDARD
	case 2:
		if args[1] == STANDARD || args[1] == SHADOW || args[1] == THINKERTOY {
			lToColorize, input, desiredFont = args[0], args[0], args[1]
		} else {
			lToColorize, input, desiredFont = args[0], args[1], STANDARD
		}
	case 3:
		lToColorize, input, desiredFont = args[0], args[1], args[2]
	default:
		err = errors.New("invalid arguments input")
		return
	}
	return
}
