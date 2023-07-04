package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func File(name string, isDstFile bool) (*os.File, error) {
	exist, err := IsExist(name)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrFileNotFound
	}
	if isDstFile {
		dst, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		return dst, nil
	}
	src, err := os.OpenFile(name, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func IsExist(FileName string) (bool, error) {
	_, err := os.Stat(FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Copy(srcFile *os.File, destFile *os.File) error {
	_, err := io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}
