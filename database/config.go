package database

import "github.com/limlance99/crreviewapi/models"

// DropAndMigrate drops existing tables then creates them again
func DropAndMigrate(models ...interface{}) {

	// comment this out if not debugging
	for _, table := range models {
		Db.Debug().DropTableIfExists(table)
		Db.Debug().AutoMigrate(table)
	}

}

// AddForeignKeys initializes the foreign keys for the tables
func AddForeignKeys() {

	// Foreign key of reviews referencing a CR
	Db.Debug().Model(&models.Review{}).AddForeignKey("crid", "crs(id)", "CASCADE", "CASCADE")

	// Foreign key of CRs referencing their location
	Db.Debug().Model(&models.CR{}).AddForeignKey("locationid", "locations(locationid)", "CASCADE", "CASCADE")

	// Foreign keys of FacilityAvailable
	Db.Debug().Model(&models.FacilityAvailable{}).AddForeignKey("crid", "crs(id)", "CASCADE", "CASCADE")
	Db.Debug().Model(&models.FacilityAvailable{}).AddForeignKey("fid", "facilities(fid)", "CASCADE", "CASCADE")

}