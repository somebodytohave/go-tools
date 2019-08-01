package mredis

import (
	"github.com/go-redis/redis"
)

// RedisClient 连接池
var RedisClient *redis.Client

// Setup 初始化连接池
func Setup(host, password string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
