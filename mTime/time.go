package mTime

import (
	"github.com/jinzhu/now"
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

// GetMonthDiff 获取当前月份的开始与结束 时间戳
func GetMonthDiff(date string) (int64, int64, error) {
	parseTime, err := now.Parse(date)
	if err != nil {
		return 0, 0, err
	}
	month := now.New(parseTime)
	beginMonth := month.BeginningOfMonth()
	lastMonth := month.EndOfMonth()
	return beginMonth.Unix(), lastMonth.Unix(), nil
}
