package model

import "github.com/jinzhu/gorm"

//Facility Facilities
type Facility struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:text; not null"`
	Description string `json:"description" gorm:"type:text"`
}

//FacilitiesManage ...
type FacilitiesManage struct {
	gorm.Model
	FacilityName string `json:"facility name" gorm:"type:text; not null"`
	Room         int    `json:"room" gorm:"type:int; not null"`
	Quantity     int    `json:"quantity" gorm:"type:int; not null; default:0"`
	Logs         string `json:"logs" gorm:"type:text"`
}
