/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	Jan 20, 2020: Lance Lim - Initialized file.
	Jan 21, 2020: Fixed foreign key associations.
*/

package models

import "time"

// CR contains the record of all CRs that have been inputted by users.
type CR struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"createdat"`
	Location   Location  `gorm:"foreignkey:LocationID;association_foreignkey:LocationID" json:"location"`
	Reviews    []Review  `gorm:"foreignkey:CRid;association_foreignkey:CRid" json:"reviews"`
	LocationID uint      `gorm:"column:locationid; not null" json:"locationid"`
	Floor      int       `gorm:"column:floor; type:int; not null" json:"floor"`
	Gender     string    `gorm:"column:gender; type:char(1); not null" json:"gender"`
}
