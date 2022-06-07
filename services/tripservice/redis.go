package main

import (
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func InitRedisConnection(uri string) error {
	opts, err := redis.ParseURL(uri)
	if err != nil {
		return err
	}

	rdb = redis.NewClient(opts)
	return nil
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		panic("need to initialize redis connection beforehand")
	}
	return rdb
}

func CloseRedisConnection() {
	if rdb != nil {
		_ = rdb.Close()
	}
}
