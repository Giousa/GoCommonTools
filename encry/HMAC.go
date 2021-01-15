/**
 *@Desc: 认证
场景：A 向 B发送了一条重要消息，生成散列函数，例如MD5加密后，将加密后消息生成认证码发过去。B 通过发来的加密消息，配合认证码和秘钥，判断这条消息的真实性
主要是判断消息的真实性，而不是用来对消息进行加密和解密
 *@Author:Giousa
 *@Date:2021/1/14
 */
package encry

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// 生成消息认证码
func GenerateHMAC(src, key []byte) string {
	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)

	fmt.Println("生成消息认证码字节数组：",result)
	fmt.Println("生成消息认证码字符串：",base64.StdEncoding.EncodeToString(result))

	return base64.StdEncoding.EncodeToString(result)
}

//验证消息认证码
func VerifyHMAC(res, src, key []byte) bool {

	// 1. 创建一个底层采用sha256算法的 hash.Hash 接口
	myHmac := hmac.New(sha256.New, key)
	// 2. 添加测试数据
	myHmac.Write(src)
	// 3. 计算结果
	result := myHmac.Sum(nil)
	// 4. 比较结果
	return hmac.Equal(res, result)
}

