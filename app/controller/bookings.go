package controller

import (
	"cp/app/model"
	"cp/app/server"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// SelectClients shows web page with filled tables of clients
func SelectBookings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookings, err := server.SelectBookings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Bookings []model.Booking
	}{
		bookings,
	}
	path := filepath.Join("public", "bookings.html")
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

// InsertClient reads client info from form and inserts it into database
func InsertBooking(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var newClient model.Client
	var newBooking model.Booking
	newClient.ClientLastName = r.FormValue("insert_second_name")
	newClient.ClientFirstName = r.FormValue("insert_name")
	newClient.ClientPhone = r.FormValue("insert_phone")
	newClient.ClientEmail = r.FormValue("insert_email")
	newClient.ClientGender = r.FormValue("insert_sex")
	if err := server.InsertNewClient(newClient); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	newBooking.ClientFirstName = r.FormValue("insert_name")
	newBooking.ClientLastName = r.FormValue("insert_second_name")
	newBooking.RoomNumber = r.FormValue("insert_room_number")
	newBooking.CheckInDate = r.FormValue("insert_check_in_date")
	newBooking.CheckOutDate = r.FormValue("insert_check_out_date")
	// newBooking.TotalPrice = r.FormValue()
	// newBooking.TotalPrice = server.GetPriceForRoomCategory()
	roomCategory, err := server.GetRoomCategory(newBooking.RoomNumber)
	if err != nil {
		return
	}
	// fmt.Println("room category ", roomCategory)

	base_price, err := server.GetPriceForRoomCategory(roomCategory)
	if err != nil {
		return
	}
	fmt.Println(newBooking.CheckInDate, newBooking.CheckOutDate)

	timeFormat := "2006-01-02"
	checkIn, err := time.Parse(timeFormat, newBooking.CheckInDate)
	if err != nil {
		fmt.Println("Error parsing CheckInDate:", err)
		return
	}

	checkOut, err := time.Parse(timeFormat, newBooking.CheckOutDate)
	if err != nil {
		fmt.Println("Error parsing CheckOutDate:", err)
		return
	}

	duration := checkOut.Sub(checkIn)

	days := int(duration.Hours() / 24)

	newBooking.TotalPrice = base_price * float64(days)

	fmt.Println("hello from insert booking")
	fmt.Println(newBooking.ClientLastName, newBooking.ClientFirstName,
		newBooking.RoomNumber,
		newBooking.CheckInDate,
		newBooking.CheckOutDate,
		newBooking.TotalPrice)

	if err := server.InsertNewBooking(newBooking); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

// DeleteClient reads client info from form and deletes client from database
func DeleteBooking(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteBooking(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
