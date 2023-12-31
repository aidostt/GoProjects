package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	//TODO:Implement notFound and methodNotAllowed http responses and assign them to the router functions
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodPost, "/ascii-art", app.postHandler)
	router.HandlerFunc(http.MethodPost, "/export", app.exportHandler)
	return router
}
