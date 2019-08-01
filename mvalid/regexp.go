package mvalid

import (
	"regexp"
)

const (
	RegexPhone    = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	RegexEmail    = "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,4})$"
	RegexUserName = "^([A-Za-z_])+\\w"
)

var (
	PhoneRegex    = regexp.MustCompile(RegexPhone)
	UsernameRegex = regexp.MustCompile(RegexUserName)
)

// 检查手机号
func RegPhone(phone string) bool {
	return PhoneRegex.MatchString(phone)
}

// 检查用户名
func RegUserName(username string) bool {
	return UsernameRegex.MatchString(username)
}

