package main

import (
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/config"
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/db"
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
