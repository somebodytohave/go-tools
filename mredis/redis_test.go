package mredis_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/sun-wenming/go-tools/mredis"
	"log"
	"testing"
)

type SetTest struct {
	Name string `json:"name"`
}

func TestRedis(t *testing.T) {

	assertions := assert.New(t)
	// 连接
	err := mredis.Setup(":6379", "")
	assertions.NoError(err)

	// key value
	//result1, err := mredis.SetKeyValue("key1", "value1", time.Second*2)
	//assertions.Equal("OK", result1)
	//assertions.NoError(err)
	//value, err := mredis.GetString("key1")
	//assertions.NoError(err)
	//assertions.Equal("value1", value)
	// 集合
	test := SetTest{Name: "1233"}
	marshal, err := json.Marshal(test)
	if err != nil {
		log.Fatalln(err)
	}
	member := string(marshal)
	log.Println(member)
	setkey := "setkey"
	err = mredis.SetSet(setkey, member, member)
	assertions.NoError(err)
	sets, err := mredis.GetSets(setkey)
	log.Println("sets..", sets)
	assertions.NoError(err)
	exist, err := mredis.ExistSetMember(setkey, member)
	assertions.NoError(err)
	log.Println("ExistSetMember..", member)
	if exist {
		err := mredis.RemoveSetMembers(setkey, member)
		assertions.NoError(err)
	}
	setss, err := mredis.GetSets(setkey)
	log.Println("sets..", setss)
	assertions.NoError(err)

}
