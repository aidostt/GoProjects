package command

import (
	"strings"
)

// Check function processes an array of strings containing commands and modifies the words based on those commands.
func Check(words *[]string) error {
	var exist bool

	// First loop: Process single-word commands.
	for i := 0; i < len(*words); i++ {
		if _, exist = Сommands[(*words)[i]]; exist {
			switch (*words)[i] {
			case "(hex)":
				(*words)[i-1] = hex((*words)[i-1]) // Apply hex transformation to the preceding word.
			case "(bin)":
				(*words)[i-1] = bin((*words)[i-1]) // Apply binary transformation to the preceding word.
			case "(up)":
				(*words)[i-1] = strings.ToUpper((*words)[i-1]) // Convert the preceding word to uppercase.
			case "(low)":
				(*words)[i-1] = strings.ToLower((*words)[i-1]) // Convert the preceding word to lowercase.
			case "(cap)":
				(*words)[i-1] = strings.Title((*words)[i-1]) // Capitalize the preceding word.
			default:
				return ErrCommNotFound // Return an error if the command is not found.
			}
			*words = delAtInd(*words, i) // Remove the command from the array.
			i--                          // Decrement the loop counter to adjust for the removed command.
		}
	}

	// Second loop: Process multi-word commands.
	var (
		num int
		err error
	)
	for i := 0; i < len(*words); i++ {
		if _, exist := advCommands[(*words)[i]]; exist {
			if i+1 > len(*words)-1 {
				return ErrInvalidInput
			}
			num, err = number((*words)[i+1]) // Get the number of words to process.
			if err != nil || i-num < 0 {
				return ErrInvalidInput
			}
			switch (*words)[i] {
			case "(up,":
				// Apply uppercase transformation to the specified range of preceding words.
				for j := i - 1; j >= i-num; j-- {
					(*words)[j] = strings.ToUpper((*words)[j])
				}
			case "(low,":
				// Apply lowercase transformation to the specified range of preceding words.
				for j := i - 1; j >= i-num; j-- {
					(*words)[j] = strings.ToLower((*words)[j])
				}
			case "(cap,":
				// Apply title case transformation to the specified range of preceding words.
				for j := i - 1; j >= i-num; j-- {
					(*words)[j] = strings.Title((*words)[j])
				}
			}
			*words = delAtInd(*words, i+1) // Remove the number and the multi-word command.
			*words = delAtInd(*words, i)
			i-- // Decrement the loop counter.
		}
	}
	return nil
}
