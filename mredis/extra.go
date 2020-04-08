package mredis

import "time"

// 通过 redis 限速, 防止连击
func LimitRate(key string, expTime time.Duration) bool {
	existKeys, _ := ExistKeys(key)
	if existKeys {
		return false
	}
	_ = SetKeyValue(key, "limit_rate", expTime)
	return true
}
