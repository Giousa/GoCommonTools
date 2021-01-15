/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/14
 */
package service

import "common-tools/encry"



func MD5(str string) string {
	return encry.GetMD5(str)
}

func Sha1(filePath string) (string,error) {
	return encry.GetSha1(filePath)
}

func HMAC(src, key []byte) string {
	return encry.GenerateHMAC(src,key)
}

func VerifyHMAC(res, src, key []byte) bool {
	return encry.VerifyHMAC(res,src,key)
}

func TripleDESEncrypt(src, key []byte) string{
	return encry.TripleDESEncrypt(src,key)
}

func TripleDESDecrypt(src, key []byte) string {
	return encry.TripleDESDecrypt(src,key)
}

func AESEncrypt(src, key []byte) string{
	return encry.AESEncrypt(src,key)
}

func AESDecrypt(src, key []byte) string{
	return encry.AESDecrypt(src,key)
}

func DesEncrypt(src, key []byte) string{
	return encry.DesEncrypt_CBC(src,key)
}

func DesDecrypt(src, key []byte) string {
	return encry.DesDecrypt_CBC(src,key)
}

func XOREncrypt(src, key int) string{

	return encry.XOREncrypt(src,key)
}

func XORDecrypt(src, key int) string{

	return encry.XORDecrypt(src,key)
}