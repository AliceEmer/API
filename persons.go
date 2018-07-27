package main

import (
	"fmt"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func (cn *Controller) allPersons() ([]*Person, error) {

	rows, err := cn.db.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pers := make([]*Person, 0)
	for rows.Next() {
		p := new(Person)
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

func (cn *Controller) personByID(id string) ([]*Person, error) {

	rows, err := cn.db.Query("SELECT * FROM person WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pers := make([]*Person, 0)
	for rows.Next() {
		p := new(Person)
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

func (cn *Controller) addPerson(p *Person) error {

	fmt.Printf("NAME: %v, %v  ", p.Firstname, p.Lastname)

	_, err := cn.db.Exec("INSERT INTO person VALUES ($1, $2)", p.Firstname, p.Lastname)
	if err != nil {
		panic(err)
	}
	return nil

}

func (cn *Controller) dropPerson(id string) error {

	_, err := cn.db.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil

}
