package model

import (
	"DormAppBackend/db"
)

// InitMigration create the database, base on the model
func InitMigration() {
	db := db.GetDB()
	// if db.HasTable(&User{}) {
	// 	db.DropTable(&User{})
	// }
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Student{}, &MoneyManage{}, &Facility{}, &FacilitiesManage{}, &Room{}, &Request{}, &Notification{})
	db.AutoMigrate()
	// defer db.Close()
}
