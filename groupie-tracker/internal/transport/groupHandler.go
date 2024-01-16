package transport

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func (t *Transport) GroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	temp := strings.TrimLeft(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(temp)
	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/groupPage.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	group, err := t.r.GetGroupById(id)
	if err != nil {
		if err.Error() == "Not Found" {
			errorHandler(w, http.StatusNotFound)
		} else {
			errorHandler(w, http.StatusInternalServerError)
		}
		return
	}

	if err := tmpl.Execute(w, group); err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
