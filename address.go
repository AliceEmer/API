package main

type Address struct {
	ID       string `json:"id,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	PersonID string `json:"person_id,omitempty"`
}

func (cn *Controller) addressByID(id string) ([]*Address, error) {

	rows, err := cn.db.Query("SELECT * FROM address WHERE person_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adds := make([]*Address, 0)
	for rows.Next() {
		add := new(Address)
		err := rows.Scan(&add.City, &add.State, &add.ID, &add.PersonID)
		if err != nil {
			return nil, err
		}
		adds = append(adds, add)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return adds, nil
}

func (cn *Controller) addAddress(a *Address, id string) error {

	_, err := cn.db.Exec("INSERT INTO address(city, state, person_id) VALUES ($1,$2,$3)", a.City, a.State, id)

	if err != nil {
		panic(err)
	}

	return nil

}

func (cn *Controller) dropAddress(id string) error {

	_, err := cn.db.Exec("DELETE FROM address WHERE person_id = $1", id)

	if err != nil {
		panic(err)
	}

	return nil

}
