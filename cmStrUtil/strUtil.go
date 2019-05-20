package cmStrUtil

func SetDefaultString(str, defaultStr string) string {
	if str == "" {
		str = defaultStr
	}
	return str
}
