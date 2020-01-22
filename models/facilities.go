package models

// Facility contains all of the CR facilities.
type Facility struct {
	Fid          uint   `gorm:"column:fid; primary_key; AUTO_INCREMENT"`
	FacilityName string `gorm:"column:facilityname; type:varchar(200); UNIQUE; not null"`
}