package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// GetCRs returns all the CRs in the database
// <url>/api/crs
func GetCRs(c echo.Context) error {
	crs := []models.CR{}

	if database.Db.Debug().Preload("Location").Find(&crs); crs == nil {
		return setErrorMessage(c)
	}

	return returnData(c, crs)
}

// GetLocations returns all the Locations in the database
// <url>/api/locations
func GetLocations(c echo.Context) error {
	locations := []models.Location{}

	if database.Db.Debug().Find(&locations); locations == nil {
		return setErrorMessage(c)
	}

	return returnData(c, locations)
}

// GetFacilities returns all the Facilities in the database
// <url>/api/facilities
func GetFacilities(c echo.Context) error {
	facilities := []models.Facility{}

	if database.Db.Debug().Find(&facilities); facilities == nil {
		return setErrorMessage(c)
	}

	return returnData(c, facilities)
}

// GetCRReviews returns all the reviews of one CR
// <url>/api/reviews/:id
func GetCRReviews(c echo.Context) error {
	cr := models.CR{}
	crid := c.Param("id")

	database.Db.Debug().Set("gorm:auto_preload", true).First(&cr, crid)
	if cr.ID == 0 {
		return setErrorMessage(c)
	}

	return returnData(c, cr)
}

// GetCRFacilities returns all the Facilities of one CR
// <url>/api/facilities/:id
func GetCRFacilities(c echo.Context) error {
	facilities := []models.Facility{}
	crid := c.Param("id")

	database.Db.Debug().Table("facilities").Select("fid, facilityname").Joins("NATURAL JOIN facilityavailable").Where("crid = ?", crid).Scan(&facilities)

	return returnData(c, facilities)
}
