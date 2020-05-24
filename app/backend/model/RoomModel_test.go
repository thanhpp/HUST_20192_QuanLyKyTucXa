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
	if err := db.GetDB().Table("room").Select("room_id").Where("room_id = ?", newRoom.RoomID).Error; err != nil {
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
