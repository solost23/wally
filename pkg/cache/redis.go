package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"wally/global"
)

func RedisConnFactory(db int) (*redis.Client, error) {
	redisConfig := global.ServerConfig.RedisConfig
	if global.RedisMapPool[db] != nil {
		return global.RedisMapPool[db], nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		DB:   db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	global.RedisMapPool[db] = rdb
	return global.RedisMapPool[db], nil
}
