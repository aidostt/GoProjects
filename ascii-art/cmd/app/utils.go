package main

import (
	"aidostt.ascii-art/pkg"
	"errors"
	"os"
	"unicode"
)

var STANDARD_HASH = "ac85e83127e49ec42487f272d9b9db8b"

type font struct {
	hash string
	path string
}

func validateInput(s string) error {
	for _, el := range s {
		if el < 32 || el > unicode.MaxASCII {
			return errors.New(pkg.ErrInvalidInput)
		}
	}
	return nil
}

func validateFiles() error {
	var fonts map[string]font
	//fill the map with libraries
	//make a loop that will iterate through fonts map, hash each data and compare them with right ones
	file, err := file(fonts["standard"].path)
	defer file.Close()
	if err != nil {
		return err
	}
	str, err := fileData(file)
	hashed, err := hash(str)
	if err != nil {
		return err
	}
	//convert to hash variables
	//compare them
	return nil
}

func fileData(file *os.File) (string, error) {
	bytes := make([]byte, 2056)
	n, err := file.Read(bytes)
	if err != nil {
		return "", err
	}
	return string(bytes[:n]), nil
}

func file(dir string) (*os.File, error) {
	exist, err := exist(dir)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrFileNotFound
	}
	file, err := os.OpenFile(dir, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func exist(FileName string) (bool, error) {
	_, err := os.Stat(FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func hash(data string) (out string, err error) {

}
