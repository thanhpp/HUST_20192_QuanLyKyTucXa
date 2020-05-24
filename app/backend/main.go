package main

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/model"
	"DormAppBackend/server"
	"DormAppBackend/tlog"
)

func main() {
	config.Init()
	db.Init()
	db.RedisClientInit("1")
	model.InitMigration()
	tlog.Init()
	server.Init()
}
