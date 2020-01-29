/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
*/

package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for postgres
	"github.com/joho/godotenv"
	"github.com/limlance99/crreviewapi/models"
)

// Db is the pointer to our database
var Db *gorm.DB

// Connect to the database
func Connect() {
	// err: variable to store any potential error (for error handling)
	var err error

	_ = godotenv.Load()

	// dbURL: URL that will be used to connect to the database, retrieved from environment variables
	dbURL := os.Getenv("DATABASE_URL")

	Db, err = gorm.Open("postgres", dbURL)

	if err != nil {
		fmt.Println(err)
	}

	// Drops existing tables then creates them again (resets db)
	DropAndMigrate(
		&models.FacilityAvailable{},
		&models.Review{},
		&models.CR{},
		&models.Location{},
		&models.Facility{},
	)

	// adding the foreign keys
	AddForeignKeys()

	// Temporary db filler function
	PopulateDB()
}
