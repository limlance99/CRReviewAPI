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
	var err error

	_ = godotenv.Load()

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
