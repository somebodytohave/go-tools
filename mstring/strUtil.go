package mstring

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 获取字符串,如果为空,设为默认值
func GetStringByDefault(str, defaultStr string) string {
	if str == "" {
		str = defaultStr
	}
	return str
}

func String(res interface{}) string {
	return fmt.Sprintf("%+v", &res)
}

func StringPrint(res interface{}) string {
	sprintf := fmt.Sprintf("%+v", &res)
	fmt.Println(sprintf)
	return sprintf
}

func StringStruct(res interface{}) string {
	//下方方法可以按行输出结构体
	b, err := json.Marshal(&res)
	if err != nil {
		return fmt.Sprintf("%+v", &res)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", &res)
	}
	return out.String()
}

func StringStructPrint(res interface{}) string {
	var resStr string
	//下方方法可以按行输出结构体
	b, err := json.Marshal(&res)
	if err != nil {
		resStr = fmt.Sprintf("%+v", &res)
		fmt.Println(resStr)
		return resStr
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		resStr = fmt.Sprintf("%+v", &res)
		fmt.Println(resStr)
		return resStr
	}
	resStr = out.String()
	fmt.Println(resStr)
	return resStr
}
