package config

import redis2 "github.com/go-redis/redis"

var redisConn *redis2.Client

func initRedis() {

	redisConn = redis2.NewClient(&redis2.Options{
		Addr:     GlobalConfig.GetString("redis.addr"),
		Password: GlobalConfig.GetString("redis.password"),
		DB:       GlobalConfig.GetInt("redis.database"),
	})

}
