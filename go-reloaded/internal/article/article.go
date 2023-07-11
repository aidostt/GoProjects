package article

import "go-reloaded.aidostt.net/internal/command"

var (
	vowels = map[rune]rune{
		'a': ' ', 'e': ' ', 'i': ' ', 'o': ' ', 'u': ' ', 'h': ' ',
		'A': ' ', 'E': ' ', 'I': ' ', 'O': ' ', 'U': ' ', 'H': ' ',
	}
)

func Check(words []string) ([]string, error) {
	for i, word := range words {
		if word == "a" || word == "A" {
			if i+1 > len(words)-1 {
				return nil, command.ErrInvalidInput
			}
			if _, exist := vowels[rune(words[i+1][0])]; exist {
				if word == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
		if word == "an" || word == "An" {
			if _, exist := vowels[rune(words[i+1][0])]; !exist {
				if word == "an" {
					words[i] = "a"
				} else {
					words[i] = "A"
				}
			}
		}
	}
	return words, nil
}
