package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for postgres
	"github.com/joho/godotenv"
)

// Db is the pointer to our database
var Db *gorm.DB

// Connect to the database
func Connect() {
	var err error

	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")

	Db, err = gorm.Open("postgres", dbURL)

	if err != nil {
		fmt.Println(err)
	}

	// comment this out if not debugging
	Db.DropTableIfExists(&FacilityAvailable{}, &Review{}, &CR{}, &Location{}, &Facility{})

	Db.AutoMigrate(
		&Location{},
		&CR{},
		&Review{},
		&Facility{},
		&FacilityAvailable{},
	)

	// adding the foreign keys
	AddForeignKeys()

	// Temporary db filler function
	PopulateDB()
}
