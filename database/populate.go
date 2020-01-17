package database

import (
	"net/http"

	"github.com/labstack/echo"
)

// PopulateDB temporary thing that populates the db
func PopulateDB() {

	facilities := []Facility{
		Facility{FacilityName: "Hand Soap"},
		Facility{FacilityName: "Bidet"},
		Facility{FacilityName: "Toilet Paper"},
	}

	for _, facility := range facilities {
		Db.Debug().Create(&facility)
	}

	locations := []Location{
		Location{Address: "Palma Hall", X: 0, Y: 0},
		Location{Address: "DCS", X: 0, Y: 0},
		Location{Address: "Math", X: 0, Y: 0},
	}

	for _, location := range locations {
		Db.Debug().Create(&location)
	}

	crs := []CR{
		CR{LocationID: 1, Floor: 1, Gender: "M"},
		CR{LocationID: 1, Floor: 1, Gender: "F"},
		CR{LocationID: 2, Floor: 2, Gender: "M"},
		CR{LocationID: 3, Floor: 1, Gender: "F"},
	}

	for _, cr := range crs {
		Db.Debug().Create(&cr)
	}

	facilitiesavailable := []FacilityAvailable{
		FacilityAvailable{CRid: 1, Fid: 1},
		FacilityAvailable{CRid: 1, Fid: 2},
		FacilityAvailable{CRid: 2, Fid: 2},
		FacilityAvailable{CRid: 4, Fid: 3},
	}

	for _, facilityavailable := range facilitiesavailable {
		Db.Debug().Create(&facilityavailable)
	}
}

type h map[string]interface{}

// TestThing is for testing
func TestThing(c echo.Context) error {
	message := &h{
		"test": "hello this works!",
	}
	return c.JSONPretty(
		http.StatusOK,
		message,
		"   ",
	)
}
