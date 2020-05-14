package db

import (
	"DormAppBackend/config"
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	config.Init()
	RedisClientInit("1")
	redis := GetRedisClient()
	err := redis.Ping().Err()
	if err != nil {
		t.Errorf("Can not connect to redis : %+v\n", err)
	}
}
