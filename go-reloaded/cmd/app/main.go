package main

import (
	"fmt"
	"go-reloaded.aidostt.net/internal/command"
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
	data, err := command.Check(src)
	//TODO: add error handling
	check(err)
	err = copyData(dst, data)
	check(err)
	fmt.Printf("file '%s' successfully copied to '%s'\n", args[0], args[1])
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
