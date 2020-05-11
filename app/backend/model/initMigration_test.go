package model

import (
	"testing"

	"DormAppBackend/config"
	"DormAppBackend/db"
)

func TestInitMigration(t *testing.T) {
	config.Init()
	db.Init()
	InitMigration()
}
