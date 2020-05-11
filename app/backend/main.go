package main

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
