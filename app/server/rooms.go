package server

import (
	"cp/app/model"
	"fmt"
	"strconv"
)

// SelectClients returns slice of all clients
func SelectRooms() ([]model.Room, error) {
	rows, err := db.Query(`
		SELECT * FROM rooms
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	var rooms []model.Room
	for rows.Next() {
		var temp model.Room
		err = rows.Scan(&temp.RoomID, &temp.RoomCategoryID,
			&temp.RoomNumber, &temp.RoomMaxGuests)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func SelectRoomCategories() ([]model.RoomCategory, error) {
	rows, err := db.Query(`
		SELECT * FROM room_categories
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	var categories []model.RoomCategory
	for rows.Next() {
		var temp model.RoomCategory
		err = rows.Scan(&temp.RoomCategoryID, &temp.RoomCategoryName,
			&temp.RoomCategoryDescription, &temp.RoomCategoryBasePrice)
		if err != nil {
			return nil, err
		}
		categories = append(categories, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return categories, nil
}

func GetRoomCategory(roomNumber string) (int, error) {
	var roomCategory string
	// query := "SELECT category_id FROM rooms WHERE number = '" + roomNumber + "'"
	query := "SELECT category_id FROM rooms WHERE number = $1"
	err := db.QueryRow(query, roomNumber).Scan(&roomCategory)
	// err := db.QueryRow(query, roomNumber).Scan(&roomCategory)

	fmt.Println(query)
	fmt.Println(roomCategory)

	// err := db.QueryRow(query, roomNumber).Scan(&roomCategory)
	if err != nil {
		// Handle errors, such as no rows in result set
		return -1, err
	}
	var result int
	result, err = strconv.Atoi(roomCategory)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func GetPriceForRoomCategory(room_category int) (float64, error) {
	var price float64

	query := "SELECT base_price FROM room_categories WHERE id = $1"
	err := db.QueryRow(query, room_category).Scan(&price)

	if err != nil {
		return -1, err
	}
	return price, nil
}
