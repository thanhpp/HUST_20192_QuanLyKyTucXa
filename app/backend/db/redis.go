package db

// import (
// 	"github.com/gin-contrib/sessions/cookie"
// )

// var (
// 	store = new(cookie.Store)
// )

// //CookieStoreInit create store to save cookie
// func CookieStoreInit() {
// 	*store = cookie.NewStore([]byte("secret"))
// }

// //GetCookieStore return store to save cookie
// func GetCookieStore() cookie.Store {
// 	return *store
// }

import (
	"DormAppBackend/config"
	"strconv"

	"github.com/go-redis/redis"
)

//RedisClient client to work with redis
var redisClient *redis.Client

//RedisClientInit connect to redis
func RedisClientInit(params ...string) {
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
