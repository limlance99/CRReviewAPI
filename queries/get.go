/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
*/

package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// GetCRs returns all the CRs in the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/crs
func GetCRs(c echo.Context) error {

	// crs: array of CR structs to store all the CRs that will be retrieved from the database
	crs := []models.CR{}

	if database.Db.Debug().Preload("Location").Find(&crs); crs == nil {
		return setErrorMessage(c)
	}

	return returnData(c, crs)
}

// GetLocations returns all the Locations in the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/locations
func GetLocations(c echo.Context) error {

	// locations: array of Location structs to store all the Locations that will be retrieved from the database
	locations := []models.Location{}

	if database.Db.Debug().Find(&locations); locations == nil {
		return setErrorMessage(c)
	}

	return returnData(c, locations)
}

// GetFacilities returns all the Facilities in the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/facilities
func GetFacilities(c echo.Context) error {

	// facilities: array of Facility structs to store all the Faciltiies that will be retrieved from the database
	facilities := []models.Facility{}

	if database.Db.Debug().Find(&facilities); facilities == nil {
		return setErrorMessage(c)
	}

	return returnData(c, facilities)
}

// GetCRReviews returns all the reviews of one CR
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/reviews/:id
func GetCRReviews(c echo.Context) error {

	// cr: CR struct to store the array of reviews of one CR that will be retrieved from the database
	cr := models.CR{}

	// crid: ID of the CR that will contain the reviews
	crid := c.Param("id")

	database.Db.Debug().Set("gorm:auto_preload", true).First(&cr, crid)
	if cr.ID == 0 {
		return setErrorMessage(c)
	}

	return returnData(c, cr)
}

// GetCRFacilities returns all the Facilities of one CR
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
// <url>/api/facilities/:id
func GetCRFacilities(c echo.Context) error {

	// facilities: array of facilties of a specific CR that will be retrieved from the database
	facilities := []models.Facility{}

	// crid: ID of the CR that will be joined with the faciltiies table to get its list of facilities
	crid := c.Param("id")

	database.Db.Debug().Table("facilities").Select("fid, facilityname").Joins("NATURAL JOIN facilityavailable").Where("crid = ?", crid).Scan(&facilities)

	return returnData(c, facilities)
}
