package controllers

import (
	"fmt"

	"github.com/AliceEmer/API2/models"
)

func (cn *Controller) allPersons() ([]*models.Person, error) {

	rows, err := cn.DB.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pers := make([]*models.Person, 0)
	for rows.Next() {
		p := new(models.Person)
		err := rows.Scan(&p.Firstname, &p.Lastname, &p.ID)
		if err != nil {
			return nil, err
		}
		pers = append(pers, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pers, nil
}

func (cn *Controller) personByID(id string) ([]*models.Person, error) {

	rows, err := cn.DB.Query("SELECT * FROM person WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pers := make([]*models.Person, 0)
	for rows.Next() {
		p := new(models.Person)
		err := rows.Scan(&p.Firstname, &p.Lastname, &p.ID)
		if err != nil {
			return nil, err
		}
		pers = append(pers, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pers, nil
}

func (cn *Controller) addPerson(p *models.Person) error {

	fmt.Printf("NAME: %v, %v  ", p.Firstname, p.Lastname)

	_, err := cn.DB.Exec("INSERT INTO person VALUES ($1, $2)", p.Firstname, p.Lastname)
	if err != nil {
		panic(err)
	}
	return nil

}

func (cn *Controller) dropPerson(id string) error {

	_, err := cn.DB.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil

}
