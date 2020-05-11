package model

import "github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/db"

// InitMigration create the database, base on the model
func InitMigration() {
	db := db.GetDB()
	if db.HasTable(&User{}) {
		db.DropTable(&User{})
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}
