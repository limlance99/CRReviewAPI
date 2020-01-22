package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// PostCRs adds a new CR record to the database
// <url>/api/crs
func PostCRs(c echo.Context) error {
	type data struct {
		locationid uint
		floor      int
		gender     string
		facilities []uint
	}

	requestBody := data{}
	newCR := models.CR{}

	c.Bind(&requestBody)

	newCR.LocationID = requestBody.locationid
	newCR.Floor = requestBody.floor
	newCR.Gender = requestBody.gender

	database.Db.Debug().Create(&newCR)
	database.Db.First(&newCR)

	newFacilityAvailable := models.FacilityAvailable{CRid: newCR.ID}
	for _, fid := range requestBody.facilities {
		newFacilityAvailable.Fid = fid
		database.Db.Create(&newFacilityAvailable)
	}

	return returnData(c, newCR)
}
