package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// PostCRs adds a new CR record to the database
// <url>/api/crs
func PostCRs(c echo.Context) error {
	newCR := models.CR{}
	c.Bind(&newCR)

	database.Db.Debug().Create(&newCR)

	return returnData(c, newCR)
}
