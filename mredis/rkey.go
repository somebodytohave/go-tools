package mredis

import (
	"errors"
	"time"
)

// 设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。
func SetKeyValue(key string, data interface{}, expTime time.Duration) error {
	result, err := RedisClient.Set(key, data, expTime).Result()
	if err != nil {
		return err
	}
	if result != "OK" {
		return errors.New(result)
	}
	return err
}

// 设置 key 的过期时间，key 过期后将不再可用.单位以秒计.
func ExpireKey(key string, expiration time.Duration) error {
	result, err := RedisClient.Expire(key, expiration).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.New("设置失败")
	}
	return err
}

// 以 UNIX 时间戳(unix timestamp)格式设置 key 的过期时间。
func ExpireAtKey(key string, tm time.Time) error {
	result, err := RedisClient.ExpireAt(key, tm).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.New("设置失败")
	}
	return err
}

// key 是否存在
func ExistKeys(keys ...string) error {
	_, err := ExistKeysRes(keys...)
	return err
}

// key 是否存在
func ExistKeysRes(keys ...string) (int64, error) {
	result, err := RedisClient.Exists(keys...).Result()
	if err != nil {
		return 0, err
	}
	if result < 1 {
		return 0, errors.New("key is not exist")
	}
	return result, nil
}

func DeleteKeys(keys ...string) error {
	result, err := RedisClient.Del(keys...).Result()
	if err != nil {
		return err
	}
	if result < 1 {
		return errors.New("can't find delete key")
	}
	return nil
}

// 以秒为单位返回 key 的剩余过期时间。
// 当 key 不存在时，返回 -2000000000 。
// 当 key 存在但没有设置剩余生存时间时，返回 -1000000000 。
// 否则，以秒为单位，返回 key 的剩余生存时间。
func TtlKey(key string) (time.Duration, error) {
	return RedisClient.TTL(key).Result()
}
