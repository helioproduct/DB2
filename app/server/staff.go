package server

import "cp/app/model"

// InsertNewClient inserts new client in database
func InsertNewStaff(newStaff model.Staff) error {
	_, err := db.Exec(`
		INSERT INTO staff 
		    (first_name, last_name,  
		    position, phone) 
		VALUES ($1, $2, $3, $4)`,
		newStaff.StaffFirstName, newStaff.StaffLastName,
		newStaff.StaffPosition, newStaff.StaffPhone)
	return err
}

// SelectClients returns slice of all clients
func SelectStaff() ([]model.Staff, error) {
	rows, err := db.Query(`
		SELECT * FROM staff
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	var staff []model.Staff
	for rows.Next() {
		var temp model.Staff
		err = rows.Scan(&temp.StaffID, &temp.StaffFirstName,
			&temp.StaffLastName, &temp.StaffPosition, &temp.StaffPhone)
		if err != nil {
			return nil, err
		}
		staff = append(staff, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return staff, nil
}

func DeleteStaff(id int) error {
	_, err := db.Exec(`
		DELETE FROM staff
		WHERE (id = $1)`,
		id)
	if err != nil {
		return err
	}
	return err
}
