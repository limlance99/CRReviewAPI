/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: JP Chanchico - Initialized file.
*/

package database

import "github.com/limlance99/crreviewapi/models"

// DropAndMigrate drops existing tables then creates them again
//   models: ...interface{} syntax means it accepts any number of parameters (with any datatype) when called and stores
//          them all in an array. main thing to be sent here should be the structs of the tables used.
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
