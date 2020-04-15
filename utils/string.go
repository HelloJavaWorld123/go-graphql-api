package utils

import "strings"

func StringEmpty(param string) bool {
	temParam := strings.TrimSpace(param)
	if temParam == "" || len(temParam) == 0 {
		return true
	}
	return false
}

func StringNotEmpty(param string) bool {
	return !StringEmpty(param)
}

//其中是否包含空的字符串
func StringAnyEmpty(params ...string) bool {
	if len(params) == 0 {
		return true
	}
	//其中包含空的字符串
	for _, param := range params {
		if StringEmpty(param) {
			return true
		}
	}
	//都不为空
	return false
}

//全部为空 有一个不为空 则返回 false
func stringAllEmpty(params ...string) bool {
	if len(params) == 0 {
		return true
	}

	//其中包含不为空的字符串
	for _, param := range params {
		if StringNotEmpty(param) {
			return false
		}
	}
	//所有的字符串都为空
	return true
}
