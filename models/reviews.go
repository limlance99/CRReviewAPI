package models

import "time"

// Review contains the reviews of all of the CRs.
type Review struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"createdat"`
	CRid       uint      `gorm:"column:crid" json:"crid"`
	Rating1    int       `gorm:"type:int" json:"rating1"`
	Rating2    int       `gorm:"type:int" json:"rating2"`
	Rating3    int       `gorm:"type:int" json:"rating3"`
	ReviewText string    `gorm:"column:reviewtext; type:varchar(200)" json:"reviewtext"`
}
