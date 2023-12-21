package controller

import (
	"cp/app/model"
	"cp/app/server"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func SelectRooms(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rooms, err := server.SelectRooms()
	roomcategories, err := server.SelectRoomCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Rooms           []model.Room
		RoomsCategories []model.RoomCategory
	}{
		rooms,
		roomcategories,
	}
	path := filepath.Join("public", "rooms.html")
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
