package models

import "time"

// Review contains the reviews of all of the CRs.
type Review struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	CRid       uint   `gorm:"column:crid"`
	Rating1    int    `gorm:"type:int"`
	Rating2    int    `gorm:"type:int"`
	Rating3    int    `gorm:"type:int"`
	ReviewText string `gorm:"column:reviewtext; type:varchar(200)"`
}
