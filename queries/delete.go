package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// DeleteCR deletes a CR Record from the database
func DeleteCR(c echo.Context) error {
	id := c.Param("id")

	CRtoDelete := models.CR{}

	database.Db.Debug().Delete(&CRtoDelete, "ID = ?", id)

	return returnData(c, CRtoDelete)
}

// DeleteReview deletes a Review from the database
func DeleteReview(c echo.Context) error {
	id := c.Param("reviewid")

	ReviewtoDelete := models.Review{}

	database.Db.Debug().Delete(&ReviewtoDelete, "ID = ?", id)

	return returnData(c, ReviewtoDelete)
}
