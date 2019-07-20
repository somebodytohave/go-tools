package mRedis

import (
	"errors"
	"time"
)

// 设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。
func SetKeyValue(key string, data interface{}, expTime time.Duration) (string, error) {
	return RedisClient.Set(key, data, expTime).Result()
}

// 设置 key 的过期时间，key 过期后将不再可用.单位以秒计.
func ExpireKey(key string, expiration time.Duration) (bool, error) {
	return RedisClient.Expire(key, expiration).Result()
}

// 以 UNIX 时间戳(unix timestamp)格式设置 key 的过期时间。
func ExpireAtKey(key string, tm time.Time) (bool, error) {
	return RedisClient.ExpireAt(key, tm).Result()
}

// key 是否存在
func ExistKeys(keys ...string) (bool, error) {
	result, err := RedisClient.Exists(keys...).Result()
	if err != nil {
		return false, err
	}
	if result < 1 {
		return false, nil
	}
	return true, nil
}

func DeleteKeys(keys ...string) (bool, error) {
	result, err := RedisClient.Del(keys...).Result()
	if err != nil {
		return false, err
	}
	if result < 1 {
		return false, errors.New("can't find delete key")
	}
	return true, nil
}
