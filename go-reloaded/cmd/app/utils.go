package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func file(name string, isDstFile bool) (*os.File, error) {
	exist, err := exist(name)
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

func copyDataTo(destFile *os.File, data []byte) error {
	//// To prevent situation where we moved our reading
	//// cursor, we adjust it to the beginning of the file
	//_, err := srcFile.Seek(0, io.SeekStart)
	//if err != nil {
	//	return err
	//}

	err := destFile.Truncate(0)
	if err != nil {
		return err
	}

	// Rewind the file cursor to the beginning
	_, err = destFile.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// Write the data to the destination file
	_, err = destFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func copyDataFrom(destFile *os.File) ([]string, error) {
	bytes := make([]byte, 2056)
	n, err := destFile.Read(bytes)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes[:n]), " "), nil
}
