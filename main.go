package main

import (
	"database/sql"
	"log"

	"github.com/AliceEmer/API2/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {

	database := InitDB("postgres://aliceecourtemer:password@localhost/persons?sslmode=disable")

	cn := &controllers.Controller{DB: database}

	// Echo instance
	e := echo.New()

	api := e.Group("/api") //group creation

	// Middleware: Use adds middleware to the chain which is run after router.
	e.Use(middleware.Logger())  //Loger defines the login interface
	e.Use(middleware.Recover()) //recovering from panics anywhere in the chain.

	api.GET("/people", cn.GetAllPersons)                  //return all people tab /api/people
	api.GET("/person/:id", cn.GetPersonByID)              //return a person depending on the ID
	api.GET("/person/address/:id", cn.GetAddressByPerson) //return a person address depending on the ID

	api.POST("/addperson", cn.CreatePerson)       //creation of a person
	api.POST("/addaddress/:id", cn.CreateAddress) //creation of an address and link it to a person depending on the ID

	api.DELETE("/deleteperson/:id", cn.DeletePerson)   //delete a person and its address depending on the ID
	api.DELETE("/deleteaddress/:id", cn.DeleteAddress) //delete an address depending on the ID

	// Start an HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}

func InitDB(dataSourceName string) *sql.DB {
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	//defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
