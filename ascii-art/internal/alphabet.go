package internal

import (
	"aidostt.ascii-art/pkg"
	"strings"
)

var ConstPath = "..\\..\\pkg\\"

func Alphabet(fontPath string) (map[rune]string, error) {
	fontPath += ".txt"
	file, err := pkg.File(ConstPath + fontPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	data, err := pkg.FileData(file)
	data = strings.ReplaceAll(data, "\r", "")
	if err != nil {
		return nil, err
	}
	asciiIndex, prevInd := ' ', 1
	alphabet := map[rune]string{}
	newLineCntr := 0
	skip := true
	for i, el := range data {
		if el == '\n' {
			if skip {
				skip = false
				continue
			}
			newLineCntr++
			if newLineCntr%8 == 0 {
				alphabet[asciiIndex] = data[prevInd+1 : i+1]
				asciiIndex++
				prevInd = i + 1
				skip = true
			}
		}
	}
	return alphabet, nil
}

func FormatOutput(alph map[rune]string, s string) string {
	//TODO:fix the colorize logic. When you trying to specify 1 letter to be colorized, function colorizes full word
	out := ""
	for i := 1; i <= 8; i++ {
		for _, letter := range s {
			newLineCounter, prevInd := 0, 0
			for j, element := range alph[letter] {
				if element == '\n' {
					newLineCounter++
					if newLineCounter == i {
						if prevInd != 0 {
							out += alph[letter][prevInd+1 : j]
						} else {
							out += alph[letter][:j]
						}
					}
					prevInd = j
				}
			}
		}
		out += "\n"
	}
	return out
}
