package mstr

// 获取字符串,如果为空,设为默认值
func GetStringByDefault(str, defaultStr string) string {
	if str == "" {
		str = defaultStr
	}
	return str
}
