package main

import (
	"aidostt.ascii-art/pkg"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var STANDARD = "standard"
var SHADOW = "shadow"
var THINKERTOY = "thinkertoy"

type font struct {
	hash string
	path string
}

func validator(input, desiredFont string, flags map[string]string) error {
	err := validateInput(input, desiredFont, flags)
	if err != nil {
		return err
	}

	err = validateFiles(desiredFont)
	if err != nil {
		return err
	}
	return nil
}

func validateFlags(flags map[string]string) error {
	var (
		ok  bool
		err error
	)
	if flags["output"] != "" {
		ok, err = pkg.Exist(flags["output"])
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("required file to output doesn't exists")
		}
	}
	if flags["reverse"] != "" {
		ok, err = pkg.Exist(flags["reverse"])
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("required file to reverse doesn't exists")
		}
	}

	switch flags["align"] {
	case "center", "left", "right", "justify", "":
		break
	default:
		return errors.New("invalid align type")
	}

	switch strings.ToLower(flags["color"]) {
	case "red", "green", "yellow", "blue", "purple", "cyan", "grey", "white":
		break
	default:
		return errors.New("invalid color type")
	}
	//TODO: validate color flag
	return nil
}

func validateInput(input, desiredFont string, flags map[string]string) error {
	err := validateFlags(flags)
	if err != nil {
		return err
	}

	for _, el := range input {
		if el < 32 || el > unicode.MaxASCII {
			return errors.New(pkg.ErrInvalidInput)
		}
	}

	switch desiredFont {
	case SHADOW, THINKERTOY, STANDARD, "":
		break
	default:
		return errors.New(pkg.ErrInvalidInput)
	}
	return nil
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
