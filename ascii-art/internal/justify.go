package internal

import (
	"errors"
	"fmt"
	"strings"
)

func Justify(justification string, text string) (string, error) {
	width, err := terminalSize()
	if err != nil {
		return "", errors.New("fatal: couldn't get terminal size")
	}
	lines := strings.Split(text, "\n")
	var justifiedLines []string
	// Process each line separately
	for _, line := range lines {
		spacesToAdd := width - len(line)
		fmt.Println(spacesToAdd)
		if spacesToAdd < 0 {
			return "", errors.New("input is too long to be justified")
		}
		justifiedLine, err := justifyText(line, spacesToAdd, width, justification)
		if err != nil {
			return "", err
		}
		justifiedLines = append(justifiedLines, justifiedLine)
	}
	// Join the lines with spaces and return
	return strings.Join(justifiedLines, "\n"), nil
}

func justifyText(text string, spacesToAdd, width int, justification string) (string, error) {
	switch justification {
	case "left":
		return text + strings.Repeat(" ", spacesToAdd), nil
	case "right":
		return strings.Repeat(" ", spacesToAdd) + text, nil
	case "center":
		leftSpaces := spacesToAdd / 2
		rightSpaces := spacesToAdd - leftSpaces
		return strings.Repeat(" ", leftSpaces) + text + strings.Repeat(" ", rightSpaces), nil
	default:
		return "", errors.New("invalid align type")
	}
}
