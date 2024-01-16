package transport

import (
	"net/http"
	"text/template"
)

type errorInfo struct {
	Status int
	Text   string
}

func errorHandler(w http.ResponseWriter, statusCode int) {
	tmpl, err := template.ParseFiles("ui/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	statusText := http.StatusText(statusCode)
	if err := tmpl.Execute(w, errorInfo{statusCode, statusText}); err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
