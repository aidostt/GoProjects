package transport

import (
	"net/http"
	"text/template"
)

func (t *Transport) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/index.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	groups, err := t.r.GetGroups()
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, groups); err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
