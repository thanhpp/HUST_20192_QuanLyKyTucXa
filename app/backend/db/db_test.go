package db

import (
	"testing"

	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/config"

	//Import the driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

func TestGetDB(t *testing.T) {
	config.Init()
	Init()
	db = GetDB()
	err := db.DB().Ping()
	if err != nil {
		t.Errorf("Something went wrong with the database : %+v", err)
	}
	db.Close()
}
