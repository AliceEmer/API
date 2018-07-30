package controllers

import (
	"net/http"

	"github.com/AliceEmer/API2/models"
	"github.com/labstack/echo"
)

func (cn *Controller) GetAllPersons(c echo.Context) error {

	rows, err := cn.DB.Query("SELECT * FROM person")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error selecting the persons",
		})
	}
	defer rows.Close()

	pers := make([]*models.Person, 0)
	for rows.Next() {
		p := new(models.Person)
		err := rows.Scan(&p.Firstname, &p.Lastname, &p.ID)
		if err != nil {
			return err
		}
		pers = append(pers, p)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	if len(pers) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person in the databse",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Person{
		"people": pers,
	})
}

func (cn *Controller) GetPersonByID(c echo.Context) error {

	id := c.Param("id")

	rows, err := cn.DB.Query("SELECT * FROM person WHERE id = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	pers := make([]*models.Person, 0)
	for rows.Next() {
		p := new(models.Person)
		err := rows.Scan(&p.Firstname, &p.Lastname, &p.ID)
		if err != nil {
			return err
		}
		pers = append(pers, p)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if len(pers) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Person{
		"people": pers,
	})
}

func (cn *Controller) CreatePerson(c echo.Context) error {

	person := models.Person{}
	if err := c.Bind(&person); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	_, err := cn.DB.Exec("INSERT INTO person VALUES ($1, $2)", person.Firstname, person.Lastname)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"firstname": person.Firstname,
		"Lastname":  person.Lastname,
	})

}

func (cn *Controller) DeletePerson(c echo.Context) error {

	id := c.Param("id")

	_, err := cn.DB.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, "Person deleted")

}
