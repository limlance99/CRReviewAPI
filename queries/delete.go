/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 23 - Initialized file.
*/

package queries

import (
	"github.com/labstack/echo"
	"github.com/limlance99/crreviewapi/database"
	"github.com/limlance99/crreviewapi/models"
)

// DeleteCR deletes a CR Record from the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func DeleteCR(c echo.Context) error {

	// id: ID of the CR to be deleted
	id := c.Param("id")

	// CRtoDelete: struct to store the CR to be deleted
	CRtoDelete := models.CR{}

	database.Db.Debug().Delete(&CRtoDelete, "ID = ?", id)

	return returnData(c, CRtoDelete)
}

// DeleteReview deletes a Review from the database
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func DeleteReview(c echo.Context) error {

	// id: ID of the review to be deleted
	id := c.Param("id")

	// ReviewtoDelete: struct to store the Review to be deleted
	ReviewtoDelete := models.Review{}

	database.Db.Debug().Delete(&ReviewtoDelete, "ID = ?", id)

	return returnData(c, ReviewtoDelete)
}
