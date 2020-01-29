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

// Facility contains all of the CR facilities.
type Facility struct {
	Fid          uint   `gorm:"column:fid; primary_key; AUTO_INCREMENT" json:"fid"`
	FacilityName string `gorm:"column:facilityname; type:varchar(200); UNIQUE; not null" json:"facilityname"`
}
