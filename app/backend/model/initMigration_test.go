package model

import (
	"DormAppBackend/tlog"
	"testing"

	"DormAppBackend/config"
	"DormAppBackend/db"
)

func TestInitMigration(t *testing.T) {
	config.Init()
	db.Init()
	tlog.Init()
	InitMigration()
}
