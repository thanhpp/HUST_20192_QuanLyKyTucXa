package db

import (
	"log"

	"DormAppBackend/config"

	//Import the driver for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init create connection to the db
func Init() {
	var err error

	dbConfig := config.GetDBConfig()

	db, err = gorm.Open(dbConfig.DBDriver, dbConfig.DBUser+":"+dbConfig.DBPass+"@tcp("+dbConfig.DBHost+":"+dbConfig.DBPort+")/"+dbConfig.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Can not connect to the database : %+v", err)
	}

	db.SingularTable(true)

}

//GetDB return a handler with the db
func GetDB() *gorm.DB {
	return db
}
