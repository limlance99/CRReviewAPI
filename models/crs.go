package models

import "time"

// CR contains the record of all CRs that have been inputted by users.
type CR struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	LocationID int    `gorm:"column:locationid; not null"`
	Floor      int    `gorm:"column:floor; type:int; not null"`
	Gender     string `gorm:"column:gender; type:char(1); not null"`
}
