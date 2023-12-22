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

// SelectTrainers shows web page with filled tables of trainers
func SelectStaff(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	staff, err := server.SelectStaff()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "staff.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "staff", staff); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertStaff(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newStaff model.Staff
	newStaff.StaffLastName = r.FormValue("insert_second_name")
	newStaff.StaffFirstName = r.FormValue("insert_name")
	newStaff.StaffPosition = r.FormValue("insert_position")
	newStaff.StaffPhone = r.FormValue("insert_phone")
	if err := server.InsertNewStaff(newStaff); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func DeleteStaff(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteStaff(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
