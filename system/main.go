package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	addr     string
	errorLog *log.Logger
	infoLog  *log.Logger
	conf     *config
}

type config struct {
	rootPath       string
	viewPath       string
	controllerPath string
	modelPath      string
}

func main() {

	fmt.Print("Main started")

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	conf := &config{}

	conf.rootPath = "/home/noname/R20"
	conf.viewPath = conf.rootPath + "/web/view"
	conf.controllerPath = conf.rootPath + "/web/controller"
	conf.modelPath = conf.rootPath + "/web/model"

	app := &application{
		":4000",
		errorLog,
		infoLog,
		conf,
	}

	srv := &http.Server{
		Addr:     app.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	errorLog.Fatal(srv.ListenAndServe())

}
