package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func FileData(file *os.File) (string, error) {
	bytes := make([]byte, 7512)
	n, err := file.Read(bytes)
	if err != nil {
		return "", err
	}
	return string(bytes[:n]), nil
}

func File(dir string) (*os.File, error) {
	exist, err := Exist(dir)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("File not found")
	}
	file, err := os.OpenFile(dir, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Exist(FileName string) (bool, error) {
	_, err := os.Stat(FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func PrintFile(file *os.File, data string) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	// Rewind the file cursor to the beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// Write the data to the destination file
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
