package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// GetCRs returns all the CRs in the database
func GetCRs(c echo.Context) error {
	crs := []models.CR{}

	if database.Db.Debug().Find(&crs); crs == nil {
		return setErrorMessage(c)
	}

	return returnData(c, crs)
}

// GetLocations returns all the Locations in the database
func GetLocations(c echo.Context) error {
	locations := []models.Location{}

	if database.Db.Debug().Find(&locations); locations == nil {
		return setErrorMessage(c)
	}

	return returnData(c, locations)
}

// GetFacilities returns all the Facilities in the database
func GetFacilities(c echo.Context) error {
	facilities := []models.Facility{}

	if database.Db.Debug().Find(&facilities); facilities == nil {
		return setErrorMessage(c)
	}

	return returnData(c, facilities)
}

// GetCRReviews returns all the reviews of one CR
func GetCRReviews(c echo.Context) error {
	reviews := []models.Review{}
	crid := c.Param("id")

	if database.Db.Debug().Where("crid = ?", crid).Find(&reviews); reviews == nil {
		return setErrorMessage(c)
	}

	return returnData(c, reviews)
}

// GetCRFacilities returns all the Facilities of one CR
func GetCRFacilities(c echo.Context) error {
	facilities := []models.Facility{}
	crid := c.Param("id")

	database.Db.Debug().Table("facilities").Select("fid, facilityname").Joins("NATURAL JOIN facilityavailable").Where("crid = ?", crid).Scan(&facilities)

	return returnData(c, facilities)
}
