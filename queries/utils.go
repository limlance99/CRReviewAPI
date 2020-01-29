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
	"net/http"

	"github.com/labstack/echo"
)

//errorMessage: struct that stores the error message that will be sent
type errorMessage struct {
	Message string `json:"message"`
}

// setErrorMessage is used to return errors when not found
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func setErrorMessage(c echo.Context) error {

	// message: errorMessage struct with a message initialized to "Nothing Found"
	message := &errorMessage{
		Message: "Nothing Found",
	}
	return c.JSONPretty(
		http.StatusNotFound,
		message,
		"   ",
	)
}

// returnData is used to return data back to the sender of the request
//	c echo.Context: this is a pointer to the entire request sent, including the address of the sender for the response. Auto-added when called in HTTP requests.
func returnData(c echo.Context, data interface{}) error {
	return c.JSONPretty(
		http.StatusOK,
		data,
		"   ",
	)
}
