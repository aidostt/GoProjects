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
	src, err := file(args[0], false)
	defer src.Close()
	check(err)
	dst, err := file(args[1], true)
	defer dst.Close()
	check(err)

	srcData, err := copyDataFrom(src)
	//TODO: add error handling
	check(err)
	data, err := validate(srcData)
	check(err)
	err = copyDataTo(dst, data)
	check(err)
	fmt.Printf("file '%s' successfully copied to '%s'\n", args[0], args[1])
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
