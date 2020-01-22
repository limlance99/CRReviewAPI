package queries

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// PostCRs adds a new CR record to the database
// <url>/api/crs
func PostCRs(c echo.Context) error {
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
