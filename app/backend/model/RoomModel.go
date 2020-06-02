package model

import (
	"DormAppBackend/db"
	"DormAppBackend/tlog"
	"errors"
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

func (r Room) GetAllRoom() ([]Room, error) {
	var returnListRoom []Room

	rows, err := db.GetDB().Table("room").Rows()
	if err != nil {
		tlog.Error("Can not get all room from db", err)
		return nil, err
	}

	for rows.Next() {
		var room Room
		db.GetDB().ScanRows(rows, &room)
		returnListRoom = append(returnListRoom, room)
	}
	return returnListRoom, nil
}

func (r Room) ChangeRoomOccupied(oldRoomID int, newRoomID int) error {
	var oldRoom, newRoom Room
	var err error

	err = db.GetDB().Table("room").Where("room_id = ?", oldRoomID).Find(&oldRoom).Error
	if err != nil {
		return err
	}

	err = db.GetDB().Table("room").Where("room_id = ?", newRoomID).Find(&newRoom).Error
	if err != nil {
		return err
	}

	if newRoom.Occupied == newRoom.Max {
		return errors.New("New room is full")
	}

	oldRoom.Occupied--
	newRoom.Occupied++

	err = db.GetDB().Save(&oldRoom).Error
	if err != nil {
		return err
	}

	err = db.GetDB().Save(&newRoom).Error
	if err != nil {
		return err
	}

	return nil
}
