package model

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"testing"
)

func TestGetStudentInfo(t *testing.T) {
	config.Init()
	db.Init()
	testStd := &Student{
		StudentID: 00000003,
		Name:      "test",
		DOB:       "test",
		Contact:   "test",
		Address:   "test",
		RoomID:    1,
		Priority:  1,
	}
	if db.GetDB().NewRecord(testStd) {
		if err := db.GetDB().Create(testStd).Error; err != nil {
			t.Errorf("Can not create user %+v", err)
		}
	}
	checkStudent, err := testStd.GetStudentInfo(testStd.StudentID)
	if (err != nil) || (checkStudent.StudentID != 00000003) {
		t.Errorf("Can not find user err : %+v\nuser : %+v", err, checkStudent)
	}
}

func TestGetFriends(t *testing.T) {
	config.Init()
	db.Init()
	var std = new(Student)
	testList, err := std.GetFriends(00000003)
	if err != nil || (len(testList) != 2) {
		t.Errorf("Err : %+v\nList : %+v", err, testList)
	}
}
