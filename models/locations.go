/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
	Jan 24, 2020: Lance Lim - Changed "x" and "y" to "Latitude" and "Longitude" respectively.
*/

package models

// Location contains the locations of all of the CRs.
type Location struct {
	LocationID uint    `gorm:"primary_key; column:locationid" json:"locationid"`
	Address    string  `gorm:"type:varchar(200); not null; UNIQUE" json:"address"`
	Latitude   float64 `gorm:"type:numeric(10,6); not null" json:"latitude"`
	Longitude  float64 `gorm:"type:numeric(10,6); not null" json:"longitude"`
}
