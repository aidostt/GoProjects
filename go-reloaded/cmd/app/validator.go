package main

import (
	"go-reloaded.aidostt.net/internal/article"
	"go-reloaded.aidostt.net/internal/command"
	"go-reloaded.aidostt.net/internal/punctuation"
	"strings"
)

func validate(srcData []string) (output []byte, err error) {
	srcData = deleteNilVal(srcData)
	data, err := article.Check(srcData)
	err = punctuation.Check(&data)
	if err != nil {
		//TODO: add error handling
		return
	}
	err = command.Check(&data)
	if err != nil {
		//TODO: add error handling
		return
	}
	return []byte(strings.Join(data, " ")), err
}
