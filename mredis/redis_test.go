package mredis_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sun-wenming/go-tools/mredis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	assertions := assert.New(t)
	// 连接
	err := mredis.Setup(":6379", "")
	assertions.NoError(err)

	// key value
	result1, err := mredis.SetKeyValue("key1", "value1", time.Second*2)
	assertions.Equal("OK", result1)
	assertions.NoError(err)
	value, err := mredis.GetString("key1")
	assertions.NoError(err)
	assertions.Equal("value1", value)
	// 集合
	err = mredis.SetSet("setkey", "value1", "value2")
	assertions.NoError(err)

}
