package mredis

import "time"

// 将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略
func SetSet(key string, members ...interface{}) error {
	return RedisClient.SAdd(key, members...).Err()
}

// 带有过期时间的 Set
func SetSetExpire(key string, expiration time.Duration, members ...interface{}) error {
	err := RedisClient.SAdd(key, members...).Err()
	if err != nil {
		return err
	}
	return ExpireKey(key, expiration)
}

// 获取集合所有值
func GetSets(key string) ([]string, error) {
	return RedisClient.SMembers(key).Result()
}

// 判断是否存在集合中
func ExistSetMember(key string, member interface{}) (bool, error) {
	return RedisClient.SIsMember(key, member).Result()
}

// 移除集合中的一个或多个成员元素，不存在的成员元素会被忽略。
func RemoveSetMembers(key string, members ...interface{}) error {
	return RedisClient.SRem(key, members...).Err()
}
