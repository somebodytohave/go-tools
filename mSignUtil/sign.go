package mSignUtil

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/sun-wenming/go-tools/mRandomUtil"
)

////将参数进行字典序排序
//var tempArray  = []string{ ticket, timeStamp, "aa"}
//sort.Strings(tempArray)
////将参数字符串拼接成一个字符串进行sha1加密
//var sha1String string = ""
//for _, v := range tempArray {
//	sha1String += v
//}
// SignSha1Str  sha1 签名
func SignSha1(sha1String string) string {
	h := sha1.New()
	h.Write([]byte(sha1String))
	return hex.EncodeToString(h.Sum([]byte("")))
}

func SignBase32(len int) string {
	return mRandomUtil.GetRandomBase32String(len)
}
