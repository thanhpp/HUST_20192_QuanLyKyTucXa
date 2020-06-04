package model

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/tlog"
	"errors"
	"testing"
)

func TestGetFacInfo(t *testing.T) {
	config.Init()
	db.Init()
	tlog.Init()
	newFac := &Facility{
		Name:        "O khoa phong",
		Description: "",
	}

	if check := db.GetDB().Table("facility").Select("name").Where("name = ?", newFac.Name).RecordNotFound(); check != true {
		tlog.Error("Error when query test unit", errors.New("record not found, creating a new one"))
		if err := db.GetDB().Table("facility").Create(newFac).Error; err != nil {
			tlog.Error("Can not create test unit", err)
			t.Error(err)
		}
	}
	checkFac, err := newFac.GetFacInfo(1)
	if checkFac.Name != "test" || err != nil {
		tlog.Error("Not valid value", err)
		t.Error(err)
	}
}
