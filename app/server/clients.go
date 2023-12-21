package server

import "cp/app/model"

// SelectClients returns slice of all clients
func SelectClients() ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT * FROM clients
		ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		var temp model.Client
		err = rows.Scan(&temp.ClientID, &temp.ClientFirstName,
			&temp.ClientLastName, &temp.ClientPhone, &temp.ClientEmail,
			&temp.ClientGender)
		if err != nil {
			return nil, err
		}
		clients = append(clients, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return clients, nil
}

func InsertNewClient(newClient model.Client) error {
	_, err := db.Exec(`
		INSERT INTO clients 
		    (last_name, first_name,  
		    phone, email, gender) 
		VALUES ($1, $2, $3, $4, $5)`,
		newClient.ClientLastName, newClient.ClientFirstName,
		newClient.ClientPhone, newClient.ClientEmail, newClient.ClientGender)
	return err
}
