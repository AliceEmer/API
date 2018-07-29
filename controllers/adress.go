package controllers

import (
	"github.com/AliceEmer/API2/models"
)

func (cn *Controller) addressByID(id string) ([]*models.Address, error) {

	rows, err := cn.DB.Query("SELECT * FROM address WHERE person_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adds := make([]*models.Address, 0)
	for rows.Next() {
		add := new(models.Address)
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

func (cn *Controller) addAddress(a *models.Address, id string) error {

	_, err := cn.DB.Exec("INSERT INTO address(city, state, person_id) VALUES ($1,$2,$3)", a.City, a.State, id)

	if err != nil {
		panic(err)
	}

	return nil

}

func (cn *Controller) dropAddress(id string) error {

	_, err := cn.DB.Exec("DELETE FROM address WHERE person_id = $1", id)

	if err != nil {
		panic(err)
	}

	return nil

}
