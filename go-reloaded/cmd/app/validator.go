package main

import (
	"os"
)

func Validator(file *os.File) error {
	var data []byte
	_, err := file.Read(data)
	if err != nil {
		return err
	}
	return nil
}
