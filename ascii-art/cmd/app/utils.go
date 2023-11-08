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

func validateInput(s string) error {
	for _, el := range s {
		if el < 32 || el > unicode.MaxASCII {
			return errors.New(pkg.ErrInvalidInput)
		}
	}
	return nil
}

func validateFiles() error {
	fonts := map[string]font{
		SHADOW:     {hash: "a49d5fcb0d5c59b2e77674aa3ab8bbb1", path: "..\\..\\pkg\\shadow.txt"},
		STANDARD:   {hash: "ac85e83127e49ec42487f272d9b9db8b", path: "..\\..\\pkg\\standard.txt"},
		THINKERTOY: {hash: "86d9947457f6a41a18cb98427e314ff8", path: "..\\..\\pkg\\thinkertoy.txt"},
	}
	for name, info := range fonts {
		file, err := pkg.File(info.path)
		defer file.Close()
		if err != nil {
			return err
		}
		data, err := pkg.FileData(file)
		hashed := pkg.Md5Hash(data)
		if err != nil {
			return err
		}
		if hashed != info.hash {
			return errors.New(fmt.Sprintf("font file %v has been overwritten or corrupted", name))
		}
	}
	return nil
}
