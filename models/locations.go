package models

// Location contains the locations of all of the CRs.
type Location struct {
	LocationID uint    `gorm:"primary_key; column:locationid"`
	Address    string  `gorm:"type:varchar(200); not null; UNIQUE"`
	X          float64 `gorm:"type:numeric(10,6); not null"`
	Y          float64 `gorm:"type:numeric(10,6); not null"`
}
