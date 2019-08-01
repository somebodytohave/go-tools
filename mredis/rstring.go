package mredis

// 获取指定 key 的值。如果 key 不存在，返回 nil 。如果key 储存的值不是字符串类型，返回一个错误。
func GetString(key string) (string, error) {
	return RedisClient.Get(key).Result()
}

