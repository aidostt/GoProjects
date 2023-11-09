package main

import (
	"aidostt.ascii-art/internal"
	"aidostt.ascii-art/pkg"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) > 2 {
		//TODO: add logger with error handling
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		return
	}
	font := STANDARD
	if len(os.Args[1:]) == 2 {
		font = os.Args[2]
	}

	err := validator(os.Args[1], font)
	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}
	alphabet, err := internal.Alphabet(font)
	if err != nil {
		fmt.Println(err)
		return
	}
	out := internal.FormatOutput(alphabet, os.Args[1])
	fmt.Println(out)
}
