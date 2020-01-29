/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
*/

package models

// FacilityAvailable relational table that contains the Facilities that are available in each CR.
type FacilityAvailable struct {
	CRid uint `gorm:"column:crid; primary_key" json:"crid"`
	Fid  uint `gorm:"column:fid; primary_key" json:"fid"`
}

// TableName this is for renaming the table so it doesn't default to facility_availables
func (FacilityAvailable) TableName() string {
	return "facilityavailable"
}
