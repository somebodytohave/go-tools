package mRedis

// 将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略
func SetSet(key string, members ...interface{}) error {
	return RedisClient.SAdd(key, members...).Err()
}

// 获取集合所有值
func GetSets(key string) ([]string, error) {
	return RedisClient.SMembers(key).Result()
}

// 判断是否存在集合中
func ExistSetMember(key string, member interface{}) (bool, error) {
	return RedisClient.SIsMember(key, member).Result()
}
