package model

import (
	"DormAppBackend/db"
	"DormAppBackend/tlog"

	"github.com/jinzhu/gorm"
)

//Facility Facilities
type Facility struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255) ;not null; unique"`
	Description string `json:"description" gorm:"type:text"`
}

//FacilitiesManage ...
type FacilitiesManage struct {
	gorm.Model
	RoomID     int    `json:"room" gorm:"type:int; not null"`
	FacilityID int    `json:"facility_name" gorm:"type:int; not null"`
	Quantity   int    `json:"quantity" gorm:"type:int; not null; default:0"`
	Default    int    `json:"default" gorm:"type:int; not null; default:0"`
	Logs       string `json:"logs" gorm:"type:text"`
}

//GetFacListByRoom Get facilities list by room id
func (facManage FacilitiesManage) GetFacListByRoom(roomID int) (*FacilitiesManage, error) {
	var returnFacManage = new(FacilitiesManage)
	var err error
	err = db.GetDB().Where("room_id = ?", roomID).Find(&returnFacManage).Error
	if err != nil {
		return nil, err
	}
	return returnFacManage, err
}

//GetFacInfo ...
func (fac Facility) GetFacInfo(FacID int) (*Facility, error) {
	var returnFacility = new(Facility)
	var err error
	err = db.GetDB().Table("facility").Where("id = ?", FacID).Find(&returnFacility).Error
	if err != nil {
		tlog.Error("Can not get facility from db", err)
		return nil, err
	}
	return returnFacility, err
}
