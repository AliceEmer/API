package main

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Controller struct {
	db *sql.DB
}

func main() {

	cn := Controller{}
	cn.db = InitDB("postgres://aliceecourtemer:password@localhost/persons?sslmode=disable")

	// Echo instance
	e := echo.New()

	api := e.Group("/api") //group creation

	// Middleware: Use adds middleware to the chain which is run after router.
	e.Use(middleware.Logger())  //Loger defines the login interface
	e.Use(middleware.Recover()) //recovering from panics anywhere in the chain.

	api.GET("/people", cn.getAllPersons)          //return all people tab /api/people
	api.GET("/person/:id", cn.getPerson)          //return a person depending on the ID
	api.GET("/person/address/:id", cn.getAddress) //return a person address depending on the ID

	api.POST("/addperson", cn.createPerson)       //creation of a person
	api.POST("/addaddress/:id", cn.createAddress) //creation of an address and link it to a person depending on the ID

	api.DELETE("/deleteperson/:id", cn.deletePerson)   //delete a person depending on the ID
	api.DELETE("/deleteaddress/:id", cn.deleteAddress) //delete an address depending on the ID

	// Start an HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}

//comment on lit result de Exec ?
//Struct vraiment utile ? creation de struc à chaque fois qu'on ne réutilise plus
