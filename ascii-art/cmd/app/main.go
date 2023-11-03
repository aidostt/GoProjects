package main

import (
	"aidostt.ascii-art/pkg"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		//TODO: add logger with error handling
		fmt.Errorf("Occured error: %v", pkg.ErrInvalidInput)
		return
	}
	arg := os.Args[1]
	err := validateInput(arg)
	if err != nil {
		fmt.Errorf("Occured error: %v", pkg.ErrInvalidInput)
		return
	}
	err = validateFiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	//validate input
	//validate ascii alphabet files
	//proceed the data handling all errors
	//output
}
