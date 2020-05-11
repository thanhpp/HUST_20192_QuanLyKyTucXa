package server

import (
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/router"
)

// Init create a server
func Init() {
	r := router.NewRouter()
	r.Run("localhost:8080")
}
