package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/limlance99/database_api/database"
)

func main() {
	database.Connect()
	defer database.Db.Close()

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	api := router.Group("/api")
	{
		api.GET("/testthing", database.TestThing)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	router.Logger.Fatal(router.Start(":" + port))
}
