package main

import (
	"aidostt.ascii-art/pkg"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
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
	fonts := map[string]font{
		"shadow":     {hash: "4d08f833a298632f3b197597c639dd1a", path: "..\\..\\pkg\\shadow.txt"},
		"standard":   {hash: "3eaca20016ebc5d69e861c786d22cf94", path: "..\\..\\pkg\\standard.txt"},
		"thinkertoy": {hash: "00d3accd1753b34bd5fba483cb2b8383", path: "..\\..\\pkg\\thinkertoy.txt"},
	}
	for name, info := range fonts {
		file, err := file(info.path)
		defer file.Close()
		if err != nil {
			return err
		}
		data, err := fileData(file)
		hashed := md5Hash(data)
		if err != nil {
			return err
		}
		if hashed != info.hash {
			return errors.New(fmt.Sprintf("font file %v has been overwritten or corrupted", name))
		}
	}
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
		return nil, errors.New("file not found")
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

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
