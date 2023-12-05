package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:         fmt.Sprintf("localhost:%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     app.errorLog,
	}
	app.infoLog.Printf("running server at %d\n", app.config.port)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
