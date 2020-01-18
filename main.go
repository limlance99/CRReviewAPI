package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/limlance99/crreviewapi/database"
)

func main() {
	database.Connect()
	defer database.Db.Close()

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	api := router.Group("/api/get")
	{
		api.GET("/crs", database.GetCRs)
		api.GET("/locations", database.GetLocations)
		api.GET("/facilities", database.GetFacilities)
		api.GET("/reviews/:id", database.GetCRReviews)
		api.GET("/facilities/:id", database.GetCRFacilities)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	router.Logger.Fatal(router.Start(":" + port))
}
