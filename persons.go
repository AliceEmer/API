package main

import (
	"fmt"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func allPersons() ([]*Person, error) {

	rows, err := db.Query("SELECT * FROM person")
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

func personByID(id string) ([]*Person, error) {

	rows, err := db.Query("SELECT * FROM person WHERE id = $1", id)
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

func addPerson(p *Person) error {

	fmt.Printf("NAME: %v, %v  ", p.Firstname, p.Lastname)

	_, err := db.Exec("INSERT INTO person VALUES ($1, $2)", p.Firstname, p.Lastname)
	if err != nil {
		panic(err)
	}
	return nil

}

func dropPerson(id string) error {

	_, err := db.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil

}
