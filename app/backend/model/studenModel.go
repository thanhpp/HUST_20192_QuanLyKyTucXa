package model

import (
	"github.com/jinzhu/gorm"
)

//Student ...
type Student struct {
	gorm.Model
	StudentID int    `json:"studenid" gorm:"type:int; not null; unique_index"`
	Name      string `json:"name" gorm:"type:text; not null"`
	DOB       string `json:"dob" gorm:"type:text"`
	Contact   string `json:"contact" gorm:"type:text; not null"`
	Address   string `json:"address" gorm:"type:text"`
	Room      int    `json:"room" gorm:"type:int; not null"`
	Priority  int    `json:"priority" gorm:"type:int; default:0"`
}
