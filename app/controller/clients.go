package controller

import (
	"cp/app/model"
	"cp/app/server"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func SelectClients(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	clients, err := server.SelectClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Clients []model.Client
	}{
		clients,
	}
	path := filepath.Join("public", "clients.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "data", data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
