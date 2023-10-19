package article

import "go-reloaded.aidostt.net/internal/command"

// Define a map that maps vowels and 'h' characters to a space to be used for word transformation.
var (
	vowels = map[rune]rune{
		'a': ' ', 'e': ' ', 'i': ' ', 'o': ' ', 'u': ' ', 'h': ' ',
		'A': ' ', 'E': ' ', 'I': ' ', 'O': ' ', 'U': ' ', 'H': ' ',
	}
)

// Check function processes an array of words and performs specific transformations.
func Check(words []string) ([]string, error) {
	for i, word := range words {
		// Check for the word "a" or "A".
		if word == "a" || word == "A" {
			if i+1 > len(words)-1 {
				// If the "a" or "A" is the last word, return an error.
				return nil, command.ErrInvalidInput
			}
			if _, exist := vowels[rune(words[i+1][0])]; exist {
				// If the next word starts with a vowel or 'h', change "a" to "an" or "A" to "An".
				if word == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
		// Check for the word "an" or "An".
		if word == "an" || word == "An" {
			if _, exist := vowels[rune(words[i+1][0])]; !exist {
				// If the next word does not start with a vowel or 'h', change "an" to "a" or "An" to "A".
				if word == "an" {
					words[i] = "a"
				} else {
					words[i] = "A"
				}
			}
		}
	}
	// Return the modified words and no error.
	return words, nil
}
