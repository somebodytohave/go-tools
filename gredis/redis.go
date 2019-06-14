package gredis

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

// RedisClient 连接池
var RedisClient *redis.Client

// Setup 初始化连接池
func Setup(host, password string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Set(key string, data interface{}, expTime time.Duration) error {
	return RedisClient.Set(key, data, expTime).Err()
}

func Exists(keys string) bool {
	result, err := RedisClient.Exists(keys).Result()
	if err != nil {
		return false
	}
	if result < 1 {
		return false
	}
	return true
}

func Get(key string) (string, error) {
	return RedisClient.Get(key).Result()
}

func Delete(keys string) (bool, error) {
	result, err := RedisClient.Del(keys).Result()
	if err != nil {
		return false, err
	}
	if result < 1 {
		return false, errors.New("can't find delete key : " + keys)
	}
	return true, nil
}
