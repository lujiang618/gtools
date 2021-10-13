package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 字符串拼接
func StringSplit(list ...string) string {
	if len(list) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, str := range list {
		builder.WriteString(str)
	}

	return builder.String()
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//首字母转大写
func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}

	list := []rune(str)
	if list[0] >= 97 && list[0] <= 122 {
		list[0] -= 32
	}

	return string(list)
}
