/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 22, 2020: Lance Lim - Initialized file.
	Jan 23, 2020: Lance Lim - Added PostReview, PostFacility, and PostLocation
*/

package queries

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// PostCR adds a new CR record to the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/crs
func PostCR(c echo.Context) error {

	// Data: This struct is so that we can get the json data from the body
	type Data struct {
		Locationid uint   `json:"locationid"`
		Floor      int    `json:"floor"`
		Gender     string `json:"gender"`
		Facilities []uint `json:"facilities"`
	}

	// requestBody: a new Data struct (for the request body to be stored in)
	requestBody := Data{}
	// newCR: a new CR struct (to be stored in the database)
	newCR := models.CR{}

	c.Bind(&requestBody)

	newCR.LocationID = requestBody.Locationid
	newCR.Floor = requestBody.Floor
	newCR.Gender = requestBody.Gender

	database.Db.Debug().Create(&newCR)

	// newFacilityAvailable: a new FacilityAvailable struct (to be stored in the database)
	newFacilityAvailable := models.FacilityAvailable{CRid: newCR.ID}

	fmt.Println(requestBody.Facilities)

	for _, fid := range requestBody.Facilities {
		newFacilityAvailable.Fid = fid
		database.Db.Debug().Create(&newFacilityAvailable)
	}

	return returnData(c, newCR)
}

// PostReview adds a new Review to the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func PostReview(c echo.Context) error {

	// newReview: a new Review Struct (to be stored in the database)
	newReview := models.Review{}
	c.Bind(&newReview)

	database.Db.Debug().Create(&newReview)

	return returnData(c, newReview)
}

// PostFacility adds a new Review to the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func PostFacility(c echo.Context) error {

	// newFacility: a new Facility struct (to be stored in the database)
	newFacility := models.Facility{}
	c.Bind(&newFacility)

	database.Db.Debug().Create(&newFacility)

	return returnData(c, newFacility)
}

// PostLocation adds a new Review to the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func PostLocation(c echo.Context) error {

	// newLocation: a new Location struct (to be stored in the database)
	newLocation := models.Location{}
	c.Bind(&newLocation)

	database.Db.Debug().Create(&newLocation)

	return returnData(c, newLocation)
}
