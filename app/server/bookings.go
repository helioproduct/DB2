package server

import "cp/app/model"

// InsertNewClient inserts new client in database
func InsertNewBooking(newBooking model.Booking) error {
	_, err := db.Exec(`
		INSERT INTO bookings 
		    (client_last_name, client_first_name,  
		    room_number, check_in_date, check_out_date, total_price) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		newBooking.ClientLastName, newBooking.ClientFirstName,
		newBooking.RoomNumber, newBooking.CheckInDate, newBooking.CheckOutDate, newBooking.TotalPrice)
	return err
}

// SelectClients returns slice of all clients
func SelectBookings() ([]model.Booking, error) {
	rows, err := db.Query(`
		SELECT * FROM bookings
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	var bookings []model.Booking
	for rows.Next() {
		var temp model.Booking
		err = rows.Scan(&temp.BookingID, &temp.ClientLastName,
			&temp.ClientFirstName, &temp.RoomNumber, &temp.CheckInDate,
			&temp.CheckOutDate, &temp.TotalPrice)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return bookings, nil
}

func DeleteBooking(id int) error {
	_, err := db.Exec(`
		DELETE FROM bookings
		WHERE (id = $1)`,
		id)
	if err != nil {
		return err
	}
	return err
}
