package controllers

import (
	"net/http"

	"github.com/AliceEmer/API2/models"
	"github.com/labstack/echo"
)

//GetAddressByPerson ... GET
func (cn *Controller) GetAddressByPerson(c echo.Context) error {

	id := c.Param("id")

	rows, err := cn.DB.Query("SELECT * FROM address WHERE person_id = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	adds := make([]*models.Address, 0)
	for rows.Next() {
		add := new(models.Address)
		err := rows.Scan(&add.City, &add.State, &add.ID, &add.PersonID)
		if err != nil {
			return err
		}
		adds = append(adds, add)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	if len(adds) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID or this person has no address",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Address{
		"address": adds,
	})
}

//CreateAddress ... POST
func (cn *Controller) CreateAddress(c echo.Context) error {

	id := c.Param("id")
	address := models.Address{}
	if err := c.Bind(&address); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	_, err := cn.DB.Exec("INSERT INTO address(city, state, person_id) VALUES ($1,$2,$3)",
		address.City, address.State, id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"city":      address.City,
		"state":     address.State,
		"person_id": id,
	})

}

//DeleteAddress ... DELETE
func (cn *Controller) DeleteAddress(c echo.Context) error {

	id := c.Param("id")

	_, err := cn.DB.Exec("DELETE FROM address WHERE person_id = $1", id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, "Address deleted")
}
