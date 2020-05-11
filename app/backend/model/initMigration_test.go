package model

import (
	"testing"

	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/config"
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/db"
)

func TestInitMigration(t *testing.T) {
	config.Init()
	db.Init()
	InitMigration()
}
