package internal

import (
	"errors"
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
		justifiedLine, ok := justifyText(line, width, justification)
		if !ok {
			return "", errors.New("invalid align option")
		}
		justifiedLines = append(justifiedLines, justifiedLine)
	}

	// Join the lines with spaces and return
	return strings.Join(justifiedLines, "\n"), nil
}

func justifyText(text string, width int, justification string) (string, bool) {
	// Calculate the number of spaces to add for justification
	spacesToAdd := width - len(text)

	// Handle different justification options
	switch justification {
	case "left":
		return text + strings.Repeat(" ", spacesToAdd), true
	case "right":
		return strings.Repeat(" ", spacesToAdd) + text, true
	case "center":
		leftSpaces := spacesToAdd / 2
		rightSpaces := spacesToAdd - leftSpaces
		return strings.Repeat(" ", leftSpaces) + text + strings.Repeat(" ", rightSpaces), true
	default:
		return "", false
	}
}
