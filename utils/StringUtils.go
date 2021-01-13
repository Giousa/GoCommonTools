/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/13
 */
package utils

import (
	"strings"
	"unicode"
)

/**
首字母小写
*/
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

/**
驼峰转蛇形：XxYy to xx-yy , XxYY to xx-y-y
*/
func SeparateToString(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '-')
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

