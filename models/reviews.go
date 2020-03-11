/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
	Mar 11, 2020: Lance Lim - Changed rating names to actual words.
*/

package models

import "time"

// Review contains the reviews of all of the CRs.
type Review struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"createdat"`
	CRid        uint      `gorm:"column:crid" json:"crid"`
	Cleanliness int       `gorm:"type:int" json:"cleanliness"`
	Facilities  int       `gorm:"type:int" json:"facilities"`
	Comfort     int       `gorm:"type:int" json:"comfort"`
	ReviewText  string    `gorm:"column:reviewtext; type:varchar(200)" json:"reviewtext"`
}
