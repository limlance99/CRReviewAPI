package database

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

	reviews := []Review{
		Review{CRid: 1, Rating1: 4, Rating2: 3, Rating3: 5, ReviewText: "Nothing special, could've been way better tbh"},
		Review{CRid: 1, Rating1: 5, Rating2: 5, Rating3: 5, ReviewText: "I love it here!"},
		Review{CRid: 2, Rating1: 4, Rating2: 3, Rating3: 5, ReviewText: "bruh ang ganda ng bidet nila dito"},
		Review{CRid: 3, Rating1: 0, Rating2: 0, Rating3: 0, ReviewText: "ew walang tissue or soap or anything >:( bad experience"},
	}

	for _, review := range reviews {
		Db.Debug().Create(&review)
	}
}
