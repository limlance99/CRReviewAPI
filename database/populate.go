package database

import "github.com/limlance99/crreviewapi/models"

// PopulateDB temporary thing that populates the db
func PopulateDB() {

	facilities := []models.Facility{
		models.Facility{FacilityName: "Hand Soap"},
		models.Facility{FacilityName: "Bidet"},
		models.Facility{FacilityName: "Toilet Paper"},
	}

	for _, facility := range facilities {
		Db.Debug().Create(&facility)
	}

	locations := []models.Location{
		models.Location{Address: "Palma Hall", X: 0, Y: 0},
		models.Location{Address: "DCS", X: 0, Y: 0},
		models.Location{Address: "Math", X: 0, Y: 0},
	}

	for _, location := range locations {
		Db.Debug().Create(&location)
	}

	crs := []models.CR{
		models.CR{LocationID: 1, Floor: 1, Gender: "M"},
		models.CR{LocationID: 1, Floor: 1, Gender: "F"},
		models.CR{LocationID: 2, Floor: 2, Gender: "M"},
		models.CR{LocationID: 3, Floor: 1, Gender: "F"},
	}

	for _, cr := range crs {
		Db.Debug().Create(&cr)
	}

	facilitiesavailable := []models.FacilityAvailable{
		models.FacilityAvailable{CRid: 1, Fid: 1},
		models.FacilityAvailable{CRid: 1, Fid: 2},
		models.FacilityAvailable{CRid: 2, Fid: 2},
		models.FacilityAvailable{CRid: 4, Fid: 3},
	}

	for _, facilityavailable := range facilitiesavailable {
		Db.Debug().Create(&facilityavailable)
	}

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
