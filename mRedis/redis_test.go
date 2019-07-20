package mRedis_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sun-wenming/go-tools/mRedis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	assertions := assert.New(t)
	// 连接
	err := mRedis.Setup(":6379", "")
	assertions.NoError(err)

	// key value
	result1, err := mRedis.SetKeyValue("key1", "value1", time.Second*2)
	assertions.Equal("OK", result1)
	assertions.NoError(err)
	value, err := mRedis.GetString("key1")
	assertions.NoError(err)
	assertions.Equal("value1", value)
	// 集合
	err = mRedis.SetSet("setkey", "value1", "value2")
	assertions.NoError(err)

}
