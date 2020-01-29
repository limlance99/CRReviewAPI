/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
	Jan 23, 2020: Lance Lim - Added routes for POST and DELETE.
*/

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

	// router: router that will be used for the API to function.
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	// api: router Group variable that allows us to set the routes for our API.
	api := router.Group("/api")
	{
		// GET requests
		api.GET("/crs", queries.GetCRs)
		api.GET("/locations", queries.GetLocations)
		api.GET("/facilities", queries.GetFacilities)
		api.GET("/reviews/:id", queries.GetCRReviews)
		api.GET("/facilities/:id", queries.GetCRFacilities)

		// POST requests
		api.POST("/crs", queries.PostCR)
		api.POST("/reviews", queries.PostReview)
		api.POST("/facilities", queries.PostFacility)
		api.POST("/locations", queries.PostLocation)

		// DELETE requests
		api.DELETE("/crs/:id", queries.DeleteCR)
		api.DELETE("/reviews/:id", queries.DeleteReview)
	}

	// port: Port that will dictate how to connect to the API. Retrieved from environment variables.
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	router.Logger.Fatal(router.Start(":" + port))
}
