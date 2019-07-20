package mRedis_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sun-wenming/go-tools/mRedis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	assertions := assert.New(t)
	err := mRedis.Setup(":6379", "")
	assertions.NoError(err)
	_, err = mRedis.SetKeyValue("key1", "value1", time.Second*2)
	time.Sleep(time.Second * 1)

	assertions.NoError(err)

	value, err := mRedis.GetString("key1")
	assertions.NoError(err)

	assertions.Equal("value1", value)
}
