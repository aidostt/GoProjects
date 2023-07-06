package main

import (
	"go-reloaded.aidostt.net/internal/article"
	"go-reloaded.aidostt.net/internal/command"
	"strings"
)

func validate(srcData []string) (output []byte, err error) {
	data, err := article.Check(srcData)
	if err != nil {
		return nil, command.ErrInvalidInput
	}
	data, err = command.Check(data)
	if err != nil {
		//TODO: add error handling
		return
	}
	return []byte(strings.Join(data, " ")), err
}
