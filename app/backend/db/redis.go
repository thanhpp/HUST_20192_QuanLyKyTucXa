package db

import (
	"DormAppBackend/config"
	"strconv"

	"github.com/go-redis/redis"
)

//RedisClient client to work with redis
var redisClient *redis.Client

func redisClientInit(params ...string) {
	rdConfig := config.GetRedisConfig()
	db, _ := strconv.Atoi(params[0])

	redisClient = redis.NewClient(&redis.Options{
		Addr:     rdConfig.Host,
		Password: rdConfig.Password,
		DB:       db,
	})
}

//GetRedisClient client for redis
func GetRedisClient() *redis.Client {
	return redisClient
}
