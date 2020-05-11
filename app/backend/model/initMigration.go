package model

import (
	"DormAppBackend/db"
)

// InitMigration create the database, base on the model
func InitMigration() {
	db := db.GetDB()
	if db.HasTable(&User{}) {
		db.DropTable(&User{})
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}
