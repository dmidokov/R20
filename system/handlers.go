package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

func (app *application) viewHander(w http.ResponseWriter, r *http.Request) {

	type Page struct {
		Title string
		Body  string
	}

	p := &Page{
		"Ptitle",
		"Pbody",
	}

	ts, err := template.ParseFiles(app.conf.rootPath + "/web/view/index.html")

	if err != nil {
		w.Write([]byte("Can't parse file"))
		return
	}

	err = ts.Execute(w, p)
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.errorLog.Output(2, trace)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
