package main

import (
	"flag"
	"github.com/go-playground/form/v4"
	"html/template"
	"log"
	"os"
)

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

type config struct {
	port int
	env  string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment, (developmen|staging|production)")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	templateCache, err := NewTemplateCache()
	formDecoder := form.NewDecoder()
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}
	err = app.serve()
	if err != nil {
		errorLog.Println(err)
		return
	}
}
