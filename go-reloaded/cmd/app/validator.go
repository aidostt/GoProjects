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
	if err != nil {
		return
	}
	err = punctuation.Check(&data)
	if err != nil {
		return
	}
	err = command.Check(&data)
	if err != nil {
		return
	}
	return []byte(strings.Join(data, " ")), err
}
