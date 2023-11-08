package main

import (
	"aidostt.ascii-art/internal"
	"aidostt.ascii-art/pkg"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		//TODO: add logger with error handling
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		return
	}
	arg := os.Args[1]
	err := validateInput(arg)
	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}
	err = validateFiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	alphabet, err := internal.Alphabet(STANDARD)
	if err != nil {
		fmt.Println(err)
		return
	}
	out := internal.FormatOutput(alphabet, os.Args[1])
	fmt.Println(out)
}
