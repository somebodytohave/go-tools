package swmTimeUtil

import (
	"time"
)

// GetMillisecond 获取毫秒的时间戳
func GetMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetTimeStamp 获取秒的时间戳
func GetTimeStamp() int64 {
	return time.Now().Unix()
}
