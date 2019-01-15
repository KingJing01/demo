package tools

import "strings"

//ParseString 解析字符串数据 A,B 解析为 'A','B'
func ParseString(data string) (result string) {
	arr := strings.Split(data, ",")
	return ParseStringArr(arr)
}

//ParseInterfaceArr 将前台的字符串数组转化 []interface{} 转 []string
func ParseInterfaceArr(params []interface{}) (param []string) {
	strArray := make([]string, len(params))
	for i, arg := range params {
		strArray[i] = arg.(string)
	}
	return strArray
}

//ParseStringArr 字符串数组转为 字符串 [A,b] 转为 'A','B'
func ParseStringArr(data []string) (result string) {
	var param string
	for _, x := range data {
		param += "'" + x + "',"
	}
	length := len(param) - 1
	params := param[0:length]
	return params
}
