package controller

import (
	"cp/app/model"
	"cp/app/server"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

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

func UpdateCategory(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_room_id"))
	if err != nil {
	 printAnswer(w, errorRes, err.Error())
	 return
	}
	category, err := strconv.Atoi(r.FormValue("update_room_category_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	number, err := strconv.Atoi(r.FormValue("update_room_number"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.UpdateCategory(id, category, number); err != nil {
	 printAnswer(w, errorRes, err.Error())
	 return
	}
	printAnswer(w, successRes, successAns)
}

func UpdatePrice(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_price_category_id"))
	if err != nil {
	 printAnswer(w, errorRes, err.Error())
	 return
	}
	price, err := strconv.ParseFloat(r.FormValue("update_price"),64)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.UpdatePrice(id, price); err != nil {
	 printAnswer(w, errorRes, err.Error())
	 return
	}
	printAnswer(w, successRes, successAns)
}
