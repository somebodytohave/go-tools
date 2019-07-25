package mGoogleauth

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var key = "MFRGGZDFMZTWQ2LKNNWG23Q="

// key 的生成 参考

//var key = mRandom.GetRandomBase32String(16)
//fmt.Println(key)
// 获取当前的时间 key 对应的 code 值
//data, e := MakeGoogleAuthenticatorForNow(key)
//fmt.Println(data)
//fmt.Println(e)
func TestMakeGoogleAuthenticator(t *testing.T) {
	Convey("TestMakeGoogleAuthenticator", t, func() {
		Convey("BeEqueal", func() {
			data, e := MakeGoogleAuthenticator(key, 1564561)
			So(e, ShouldBeNil)
			So(data, ShouldEqual, "966012")
		})
	})
}
