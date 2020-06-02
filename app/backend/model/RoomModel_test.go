package model

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/tlog"
	"testing"
)

var roomMod = new(Room)

func TestGetRoomInfo(t *testing.T) {
	config.Init()
	db.Init()
	tlog.Init()
	newRoom := &Room{
		RoomID:      1,
		Price:       1000000,
		Occupied:    3,
		Max:         8,
		Description: "",
	}
	if check := db.GetDB().Table("room").Select("room_id").Where("room_id = ?", newRoom.RoomID).RecordNotFound(); check != true {
		if err := db.GetDB().Create(newRoom).Error; err != nil {
			t.Errorf("Can not create room %+v", err)
		}
	}
	room, err := roomMod.GetRoomInfo(1)
	if err != nil || room.Price != 1000000 {
		t.Errorf("%+v", err)
	} else {
		tlog.GetLogger().Info("room get suc")
	}
}

func TestChangeRoomOccupied(t *testing.T) {
	config.Init()
	db.Init()
	tlog.Init()
	r := &Room{
		RoomID:      2,
		Price:       500000,
		Occupied:    0,
		Max:         8,
		Description: "",
	}

	check := db.GetDB().Table("room").Where("room_id = ?", r.RoomID).RecordNotFound()
	if check == true {
		err := db.GetDB().Create(r).Error
		if err != nil {
			t.Errorf("Can not create new room %+v", err)
			return
		}
	}

	err := r.ChangeRoomOccupied(2, 1)
	if err != nil {
		t.Errorf("Can not change room occupied %+v", err)
	}
}
