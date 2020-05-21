package model

import (
	"DormAppBackend/db"
	"errors"
	"fmt"
)

//Student ...
type Student struct {
	DefaultModel
	StudentID int    `json:"studenid" gorm:"type:int; not null; unique_index; primary_key"`
	Name      string `json:"name" gorm:"type:text; not null"`
	DOB       string `json:"dob" gorm:"type:text"`
	Contact   string `json:"contact" gorm:"type:text; not null"`
	Address   string `json:"address" gorm:"type:text"`
	RoomID    int    `json:"room" gorm:"type:int; not null"`
	Priority  int    `json:"priority" gorm:"type:int; default:0"`
}

//MoneyManage ...
type MoneyManage struct {
	StudentID   int ``
	Month       int
	Status      string
	Description string
}

//GetStudentInfo based on student id
func (s Student) GetStudentInfo(studentid int) (*Student, error) {
	std := new(Student)
	err := db.GetDB().Where("student_id = ?", studentid).Find(&std).Error
	if err != nil {
		return nil, errors.New("Can not find student with given id")
	}
	return std, nil
}

//GetFriends ...
func (s Student) GetFriends(studentid int) ([]Student, error) {
	var listStd []Student
	var err error
	var roomID int
	db.GetDB().Table("student").Select("room_id").Where("student_id = ?", studentid).Row().Scan(&roomID)

	rows, err := db.GetDB().Table("student").Where("room_id = ?", &roomID).Not("student_id", studentid).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var std Student
		rows.Scan(&std)
		listStd = append(listStd, std)
	}
	fmt.Printf("%+v", listStd)
	return listStd, err
}
