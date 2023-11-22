package internal

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func terminalSize() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		return 0, err
	}
	return width, nil
}
