package main

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/model"
	"DormAppBackend/server"
)

func main() {
	config.Init()
	db.Init()
	db.RedisClientInit("1")
	model.InitMigration()
	server.Init()
}
