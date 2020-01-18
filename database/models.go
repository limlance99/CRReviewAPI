package database

import (
	"time"
)

type (
	// Location contains the locations of all of the CRs.
	Location struct {
		LocationID uint    `gorm:"primary_key; column:locationid"`
		Address    string  `gorm:"type:varchar(200); not null; UNIQUE"`
		X          float64 `gorm:"type:numeric(10,6); not null"`
		Y          float64 `gorm:"type:numeric(10,6); not null"`
	}

	// Review contains the reviews of all of the CRs.
	Review struct {
		ID         uint `gorm:"primary_key"`
		CreatedAt  time.Time
		CRid       uint   `gorm:"column:crid"`
		Rating1    int    `gorm:"type:int"`
		Rating2    int    `gorm:"type:int"`
		Rating3    int    `gorm:"type:int"`
		ReviewText string `gorm:"column:reviewtext; type:varchar(200)"`
	}

	// CR contains the record of all CRs that have been inputted by users.
	CR struct {
		ID         uint `gorm:"primary_key"`
		CreatedAt  time.Time
		LocationID int    `gorm:"column:locationid; not null"`
		Floor      int    `gorm:"column:floor; type:int; not null"`
		Gender     string `gorm:"column:gender; type:char(1); not null"`
	}

	// Facility contains all of the CR facilities.
	Facility struct {
		Fid          uint   `gorm:"column:fid; primary_key; AUTO_INCREMENT"`
		FacilityName string `gorm:"column:facilityname; type:varchar(200); UNIQUE; not null"`
	}

	// FacilityAvailable relational table that contains the Facilities that are available in each CR.
	FacilityAvailable struct {
		CRid uint `gorm:"column:crid; primary_key"`
		Fid  uint `gorm:"column:fid; primary_key"`
	}
)

// TableName this is for renaming the table so it doesn't default to facility_availables
func (FacilityAvailable) TableName() string {
	return "facilityavailable"
}

// AddForeignKeys initializes the foreign keys for the tables
func AddForeignKeys() {

	// Foreign key of reviews referencing a CR
	Db.Model(&Review{}).AddForeignKey("crid", "crs(id)", "CASCADE", "CASCADE")

	// Foreign key of CRs referencing their location
	Db.Model(&CR{}).AddForeignKey("locationid", "locations(locationid)", "CASCADE", "CASCADE")

	// Foreign keys of FacilityAvailable
	Db.Model(&FacilityAvailable{}).AddForeignKey("crid", "crs(id)", "CASCADE", "CASCADE")
	Db.Model(&FacilityAvailable{}).AddForeignKey("fid", "facilities(fid)", "CASCADE", "CASCADE")

}
