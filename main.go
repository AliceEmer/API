package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	InitDB("postgres://aliceecourtemer:Chiberoua.13ass@localhost/persons?sslmode=disable")

	// Echo instance
	e := echo.New()

	api := e.Group("/api") //group creation

	// Middleware: Use adds middleware to the chain which is run after router.
	e.Use(middleware.Logger())  //Loger defines the login interface
	e.Use(middleware.Recover()) //recovering from panics anywhere in the chain.

	api.GET("/people", getAllPersons)           //return all people tab /api/people
	api.GET("/person/:id", getPerson)           //return a person depending on the ID
	api.GET("/person/:id/address/", getAddress) //return a person address depending on the ID

	api.POST("/addperson", createPerson)       //creation of a person
	api.POST("/addaddress/:id", createAddress) //creation of an address and link it to a person depending on the ID

	api.DELETE("/deleteperson/:id", deletePerson)   //delete a person depending on the ID
	api.DELETE("/deleteaddress/:id", deleteAddress) //delete an address depending on the ID

	// Start an HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}

//comment on lit result de Exec ?
//Struct vraiment utile ? creation de struc à chaque fois qu'on ne réutilise plus
