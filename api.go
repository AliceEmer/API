package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GET
func getAllPersons(c echo.Context) error {
	pers, err := allPersons()

	if err != nil || len(pers) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person in the table",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*Person{
		"people": pers,
	})
}

func getPerson(c echo.Context) error {

	id := c.Param("id")
	pers, err := personByID(id)

	if err != nil || len(pers) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*Person{
		"people": pers,
	})

}

func getAddress(c echo.Context) error {

	id := c.Param("id")

	fmt.Printf("id : %v", id)

	adds, err := addressByID(id)

	if err != nil || len(adds) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID or this person has no address",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*Address{
		"address": adds,
	})

}

//POST
func createPerson(c echo.Context) error {
	person := Person{}
	if err := c.Bind(&person); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	addPerson(&person)
	return c.JSON(http.StatusOK, map[string]string{
		"firstname": person.Firstname,
		"Lastname":  person.Lastname,
	})
}

func createAddress(c echo.Context) error {

	id := c.Param("id")
	address := Address{}
	if err := c.Bind(&address); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	addAddress(&address, id)
	return c.JSON(http.StatusOK, map[string]string{
		"city":      address.City,
		"state":     address.State,
		"person_id": id,
	})
}

//DELETE
func deletePerson(c echo.Context) error {
	id := c.Param("id")
	dropPerson(id)
	return c.JSON(http.StatusOK, "Person deleted")
}

func deleteAddress(c echo.Context) error {
	id := c.Param("id")
	dropAddress(id)
	return c.JSON(http.StatusOK, "Address deleted")
}

//comment on lit result de Exec ?
//Struct vraiment utile ? creation de struc à chaque fois qu'on ne réutilise plus
