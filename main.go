package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/queries"
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
		api.GET("/crs", queries.GetCRs)
		api.GET("/locations", queries.GetLocations)
		api.GET("/facilities", queries.GetFacilities)
		api.GET("/reviews/:id", queries.GetCRReviews)
		api.GET("/facilities/:id", queries.GetCRFacilities)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	router.Logger.Fatal(router.Start(":" + port))
}
