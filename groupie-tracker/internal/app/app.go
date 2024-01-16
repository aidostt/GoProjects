package app

import (
	"fmt"
	"groupie-tracker/internal/api"
	"groupie-tracker/internal/transport"
	"log"
	"net/http"
	"strconv"
)

type Handlers interface {
	HomeHandler(http.ResponseWriter, *http.Request)
	GroupHandler(http.ResponseWriter, *http.Request)
}

type Server struct {
	Addr     string
	routes   http.Handler
	handlers Handlers
}

func NewServer(host string, port int) *Server {
	s := &Server{
		Addr: host + ":" + strconv.Itoa(port),
	}
	apishka := api.New()
	handlers := transport.New(apishka)
	s.handlers = handlers
	s.setRoutes()
	return s
}

func (s *Server) Start() {
	fmt.Printf("Server staring at http://%v\n", s.Addr)
	if err := http.ListenAndServe(s.Addr, s.routes); err != nil {
		log.Fatal(err)
	}
}
