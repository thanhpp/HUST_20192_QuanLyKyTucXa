package model

import (
	"DormAppBackend/db"
	"DormAppBackend/tlog"
)

//Room ...
type Room struct {
	DefaultModel
	RoomID      int    `gorm:"type:int; size:10; not null; unique_index; primary_key" json:"roomID"`
	Price       int    `gorm:"type:int; size:10; not null; default:0" json:"roomPrice"`
	Occupied    int    `gorm:"type:int; size:10; not null;  default:0" json:"occupied"`
	Max         int    `gorm:"type:int; size:10; not null; default:0" json:"roomMax"`
	Description string `gorm:"type:text" json:"description"`
}

//GetRoomInfo ...
func (r Room) GetRoomInfo(roomID int) (*Room, error) {
	var returnRoom = new(Room)
	var err error
	err = db.GetDB().Where("room_id = ?", roomID).Find(&returnRoom).Error
	if err != nil {
		tlog.Error("Can't get room from db", err)
		return nil, err
	}

	return returnRoom, err
}
