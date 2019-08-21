package mredis

import "time"

// 通过 redis 限速, 防止连击
func LimitRate(key string, times int) bool {
	existKeys, _ := ExistKeys(key)
	if existKeys {
		return false
	}
	_, _ = SetKeyValue(key, "limit_rate", time.Duration(times))
	return true
}
