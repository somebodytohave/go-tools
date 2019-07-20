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

// 以秒为单位返回 key 的剩余过期时间。
// 当 key 不存在时，返回 -2000000000 。
// 当 key 存在但没有设置剩余生存时间时，返回 -1000000000 。
// 否则，以秒为单位，返回 key 的剩余生存时间。
func TtlKey(key string) (time.Duration, error) {
	return RedisClient.TTL(key).Result()
}
