package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	src, err := File(args[0], false)
	defer src.Close()
	check(err)
	dst, err := File(args[1], true)
	defer dst.Close()
	check(err)
	Validator(src)
	err = Copy(src, dst)
	check(err)
	fmt.Printf("File '%s' successfully copied to '%s'\n", args[0], args[1])

}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

//Simply add 42 (hex) and 10 "(bin)" and you will see the result "is" "68."
//contains exStr
