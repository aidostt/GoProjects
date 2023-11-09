package main

import (
	"aidostt.ascii-art/pkg"
	"errors"
	"fmt"
	"unicode"
)

var STANDARD = "standard"
var SHADOW = "shadow"
var THINKERTOY = "thinkertoy"

type font struct {
	hash string
	path string
}

func validator(s, desiredFont string) error {
	err := validateInput(s, desiredFont)
	if err != nil {
		return err
	}

	err = validateFiles(desiredFont)
	if err != nil {
		return err
	}
	return nil
}

func validateInput(s, desiredFont string) error {
	for _, el := range s {
		if el < 32 || el > unicode.MaxASCII {
			return errors.New(pkg.ErrInvalidInput)
		}
	}

	switch desiredFont {
	case SHADOW:
		return nil
	case STANDARD:
		return nil
	case THINKERTOY:
		return nil
	case "":
		return nil
	default:
		return errors.New(pkg.ErrInvalidInput)
	}
}

func validateFiles(desiredFont string) error {
	f := font{}
	switch desiredFont {
	case SHADOW:
		f = font{hash: "a49d5fcb0d5c59b2e77674aa3ab8bbb1", path: "..\\..\\pkg\\shadow.txt"}
	case THINKERTOY:
		f = font{hash: "86d9947457f6a41a18cb98427e314ff8", path: "..\\..\\pkg\\thinkertoy.txt"}
	default:
		f = font{hash: "ac85e83127e49ec42487f272d9b9db8b", path: "..\\..\\pkg\\standard.txt"}
	}

	file, err := pkg.File(f.path)
	defer file.Close()
	if err != nil {
		return err
	}
	data, err := pkg.FileData(file)
	hashed := pkg.Md5Hash(data)
	if err != nil {
		return err
	}
	if hashed != f.hash {
		return errors.New(fmt.Sprintf("%v file has been overwritten or corrupted", desiredFont))
	}

	return nil
}
