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
