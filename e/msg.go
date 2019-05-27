package e

const (
	SUCCESS  = 200
	ERROR400 = 400
	ERROR500 = 500
)

var MsgFlags = map[int]string{
	SUCCESS:  "ok",
	ERROR400: "请求参数有误",
	ERROR500: "服务器错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR500]
}
