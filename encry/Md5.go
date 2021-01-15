/**
 *@Desc: 文件、字符串等加密
 *@Author:Giousa
 *@Date:2021/1/14
 */
package encry

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5(str string) string {
	// 1. 计算数据的md5
	result := md5.Sum([]byte(str))
	// 2. 数据格式化为16进制格式字符串
	res := fmt.Sprintf("%x", result)
	// --- 这是另外一种格式化切片的方式
	res = hex.EncodeToString(result[:])

	fmt.Println("md5 原数据：",str)
	fmt.Println("md5 加密后：",res)
	return  res
}

func GetMD5Hash(str []byte) string {
	// 1. 创建一个使用MD5校验的Hash对象`
	myHash := md5.New()
	// 2. 通过io操作将数据写入hash对象中
	//io.WriteString(myHash, string(str))
	//或者
	myHash.Write(str)
	// 3. 计算结果
	result := myHash.Sum(nil)
	// 4. 将结果转换为16进制格式字符串
	res := fmt.Sprintf("%x", result)
	// --- 这是另外一种格式化切片的方式
	//res = hex.EncodeToString(result)

	return res
}
