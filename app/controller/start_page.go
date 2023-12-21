package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

// StartPage shows main web page
func StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "start_page.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
