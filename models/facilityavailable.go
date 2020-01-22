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
