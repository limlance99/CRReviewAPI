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
