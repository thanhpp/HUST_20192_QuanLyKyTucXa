package model

import (
	"DormAppBackend/db"
	"DormAppBackend/forms"
	"DormAppBackend/tlog"
	"errors"

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

func (fac Facility) NewFacility(newFacForm forms.NewFacilityForm) (*Facility, error) {
	newFac := &Facility{
		Name:        newFacForm.Name,
		Description: newFacForm.Description,
	}

	if !db.GetDB().Debug().First(&newFac, "name LIKE ?", newFac.Name).RecordNotFound() {
		return nil, errors.New("record existed")
	}

	err := db.GetDB().Create(&newFac).Error
	if err != nil {
		return nil, err
	}

	return newFac, nil
}

//GetFacListByRoom Get facilities list by room id
func (facManage FacilitiesManage) GetFacListByRoom(roomID int) (*[]FacilitiesManage, error) {
	var returnFacManage []FacilitiesManage
	var err error
	rows, err := db.GetDB().Table("facilities_manage").Where("room_id = ?", roomID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var facMan FacilitiesManage
		db.GetDB().ScanRows(rows, &facMan)

		returnFacManage = append(returnFacManage, facMan)
	}

	return &returnFacManage, err
}

func (facManage FacilitiesManage) NewFacilityManage(newFacMngForm forms.NewFacilityManageForm) (*FacilitiesManage, error) {
	facMng := &FacilitiesManage{
		RoomID:     newFacMngForm.RoomID,
		FacilityID: newFacMngForm.FacilityID,
		Quantity:   newFacMngForm.Default,
		Default:    newFacMngForm.Default,
	}

	if db.GetDB().Debug().Table("room").Select("room_id").Find(&Room{RoomID: facMng.RoomID}).RecordNotFound() {
		return nil, errors.New("No room available")
	}

	if db.GetDB().Debug().Table("facility").Select("id").Find(&Facility{Model: gorm.Model{ID: uint(facMng.FacilityID)}}).RecordNotFound() {
		return nil, errors.New("No facility available")
	}

	if !db.GetDB().Debug().First(&facMng, "room_id = ? AND facility_id = ?", facMng.RoomID, facMng.FacilityID).RecordNotFound() {
		return nil, errors.New("record existed")
	}

	err := db.GetDB().Table("facilities_manage").Create(&facMng).Error
	if err != nil {
		return nil, err
	}

	return facMng, nil
}

func (facManage FacilitiesManage) UpdateFacilityManage(facMngID int, quantity int) (*FacilitiesManage, error) {
	var facMng FacilitiesManage
	err := db.GetDB().Debug().Table("facilities_manage").Where("id = ?", facMngID).Find(&facMng).Error
	if err != nil {
		return nil, err
	}

	facMng.Quantity = quantity

	err = db.GetDB().Table("facilities_manage").Save(&facMng).Error
	if err != nil {
		return nil, err
	}

	return &facMng, nil
}
