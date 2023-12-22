package main

import (
	"ascii-art-web.aidostt.net/internal"
	"fmt"
	"net/http"
)

type asciiForm struct {
	Input string `form:"input"`
	Font  string `form:"font"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.tmpl", nil)
	//TODO:decode the form into variables

}
func (app *application) PostFormHandler(w http.ResponseWriter, r *http.Request) {
	var form asciiForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	//TODO: add limiter for the length of the input
	err = internal.Validator(form.Input, form.Font, nil)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	alphabet, err := internal.Alphabet(form.Font)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	output := internal.FormatOutput(alphabet, form.Input)
	fmt.Println(output)
	app.render(w, http.StatusOK, "view.tmpl", output)

}

func (app *application) ExportFileHandler(w http.ResponseWriter, r *http.Request) {

}
