package app

import (
	"net/http"
)

func (s *Server) setRoutes() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	mux.HandleFunc("/", s.handlers.HomeHandler)
	mux.HandleFunc("/artist/", s.handlers.GroupHandler)

	s.routes = mux
}
