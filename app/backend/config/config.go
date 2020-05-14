package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper

	dbConfig *DBConfig
	rdConfig = new(RDConfig)
)

//Init takes the environment starts the viper
func Init() {
	configInit()
	dbConfigInit()
	redisConfigInit()
}

//GetConfig return the Viper to read config from file
func configInit() {
	config = viper.New()
	config.SetConfigName("develop")
	config.SetConfigType("toml")

	config.AddConfigPath(".")
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func getConfig() *viper.Viper {
	return config
}

//DBConfig contains database config
type DBConfig struct {
	DBDriver string
	DBUser   string
	DBHost   string
	DBPort   string
	DBPass   string
	DBName   string
}

func dbConfigInit() {
	c := getConfig()
	dbConfig = &DBConfig{
		DBDriver: c.GetString("DBDriver"),
		DBUser:   c.GetString("DBUser"),
		DBHost:   c.GetString("DBHost"),
		DBPort:   c.GetString("DBPort"),
		DBPass:   c.GetString("DBPass"),
		DBName:   c.GetString("DBName"),
	}
}

//GetDBConfig return object contains config for the database
func GetDBConfig() *DBConfig {
	return dbConfig
}

//RDConfig for redis
type RDConfig struct {
	Host     string
	Password string
}

func redisConfigInit() {
	c := getConfig()
	rdConfig = &RDConfig{
		Host:     c.GetString("RedisHost"),
		Password: c.GetString("RedisPassword"),
	}
}

//GetRedisConfig config object for redis
func GetRedisConfig() *RDConfig {
	return rdConfig
}
