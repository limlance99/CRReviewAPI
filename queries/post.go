package queries

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// PostCR adds a new CR record to the database
// <url>/api/crs
func PostCR(c echo.Context) error {
	type Data struct {
		Locationid uint   `json:"locationid"`
		Floor      int    `json:"floor"`
		Gender     string `json:"gender"`
		Facilities []uint `json:"facilities"`
	}

	requestBody := Data{}
	newCR := models.CR{}

	c.Bind(&requestBody)

	newCR.LocationID = requestBody.Locationid
	newCR.Floor = requestBody.Floor
	newCR.Gender = requestBody.Gender

	database.Db.Debug().Create(&newCR)

	newFacilityAvailable := models.FacilityAvailable{CRid: newCR.ID}
	fmt.Println(requestBody.Facilities)

	for _, fid := range requestBody.Facilities {
		newFacilityAvailable.Fid = fid
		database.Db.Debug().Create(&newFacilityAvailable)
	}

	return returnData(c, newCR)
}

// PostReview adds a new Review to the database
func PostReview(c echo.Context) error {
	newReview := models.Review{}
	c.Bind(&newReview)

	database.Db.Debug().Create(&newReview)

	return returnData(c, newReview)
}
