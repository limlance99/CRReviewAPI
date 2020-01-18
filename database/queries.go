package database

import (
	"net/http"

	"github.com/labstack/echo"
)

type errorMessage struct {
	Message string `json:"message"`
}

// setErrorMessage is used to return errors when not found
func setErrorMessage(c echo.Context) error {
	message := &errorMessage{
		Message: "Nothing Found",
	}
	return c.JSONPretty(
		http.StatusNotFound,
		message,
		"   ",
	)
}

// returnData back to the request
func returnData(c echo.Context, data interface{}) error {
	return c.JSONPretty(
		http.StatusOK,
		data,
		"   ",
	)
}

// GetCRs returns all the CRs in the database
func GetCRs(c echo.Context) error {
	crs := []CR{}

	if Db.Debug().Find(&crs); crs == nil {
		return setErrorMessage(c)
	}

	return returnData(c, crs)
}

// GetLocations returns all the Locations in the database
func GetLocations(c echo.Context) error {
	locations := []Location{}

	if Db.Debug().Find(&locations); locations == nil {
		return setErrorMessage(c)
	}

	return returnData(c, locations)
}

// GetFacilities returns all the Facilities in the database
func GetFacilities(c echo.Context) error {
	facilities := []Facility{}

	if Db.Debug().Find(&facilities); facilities == nil {
		return setErrorMessage(c)
	}

	return returnData(c, facilities)
}

// GetCRReviews returns all the reviews of one CR
func GetCRReviews(c echo.Context) error {
	reviews := []Review{}
	crid := c.Param("id")

	if Db.Debug().Where("crid = ?", crid).Find(&reviews); reviews == nil {
		return setErrorMessage(c)
	}

	return returnData(c, reviews)
}

// GetCRFacilities returns all the Facilities of one CR
func GetCRFacilities(c echo.Context) error {
	facilities := []Facility{}
	crid := c.Param("id")

	Db.Debug().Table("facilities").Select("fid, facilityname").Joins("NATURAL JOIN facilityavailable").Where("crid = ?", crid).Scan(&facilities)

	return returnData(c, facilities)
}
