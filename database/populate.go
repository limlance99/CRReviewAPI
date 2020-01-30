/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: JP Chanchico - Initialized file.
	Jan 23, 2020: JP Chanchico - Updated dummy data.
*/

package database

import "github.com/limlance99/crreviewapi/models"

// PopulateDB temporary thing that populates the db
func PopulateDB() {

	// facilities: array of Facility structs holding dummy data
	facilities := []models.Facility{
		models.Facility{FacilityName: "Hand Soap"},
		models.Facility{FacilityName: "Bidet"},
		models.Facility{FacilityName: "Toilet Paper"},
	}

	for _, facility := range facilities {
		Db.Debug().Create(&facility)
	}

	// locations: array of Location structs holding dummy data
	locations := []models.Location{
		models.Location{Address: "Palma Hall", Latitude: 14.6537, Longitude: 121.0699},
		models.Location{Address: "DCS", Latitude: 14.6486, Longitude: 121.0685},
		models.Location{Address: "Math", Latitude: 14.6845, Longitude: 121.0715},
		models.Location{Address: "CAL", Latitude: 14.6527, Longitude: 121.0674},
	}

	for _, location := range locations {
		Db.Debug().Create(&location)
	}

	// crs: array of CR structs holding dummy data
	crs := []models.CR{
		models.CR{LocationID: 1, Floor: 1, Gender: "M"},
		models.CR{LocationID: 1, Floor: 1, Gender: "F"},
		models.CR{LocationID: 2, Floor: 2, Gender: "M"},
		models.CR{LocationID: 3, Floor: 1, Gender: "F"},
	}

	for _, cr := range crs {
		Db.Debug().Create(&cr)
	}

	// facilitiesavailable: array of FacilityAvailable structs holding dummy data
	facilitiesavailable := []models.FacilityAvailable{
		models.FacilityAvailable{CRid: 1, Fid: 1},
		models.FacilityAvailable{CRid: 1, Fid: 2},
		models.FacilityAvailable{CRid: 2, Fid: 2},
		models.FacilityAvailable{CRid: 4, Fid: 3},
	}

	for _, facilityavailable := range facilitiesavailable {
		Db.Debug().Create(&facilityavailable)
	}

	// reviews: array of Review structs holding dummy data
	reviews := []models.Review{
		models.Review{CRid: 1, Rating1: 4, Rating2: 3, Rating3: 5, ReviewText: "Nothing special, could've been way better tbh"},
		models.Review{CRid: 1, Rating1: 5, Rating2: 5, Rating3: 5, ReviewText: "I love it here!"},
		models.Review{CRid: 2, Rating1: 4, Rating2: 3, Rating3: 5, ReviewText: "bruh ang ganda ng bidet nila dito"},
		models.Review{CRid: 3, Rating1: 0, Rating2: 0, Rating3: 0, ReviewText: "ew walang tissue or soap or anything >:( bad experience"},
	}

	for _, review := range reviews {
		Db.Debug().Create(&review)
	}
}
