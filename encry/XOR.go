/**
 *@Desc: 按位异或
 *@Author:Giousa
 *@Date:2021/1/14
 */
package encry

func XOREncrypt(src, key int) string{

	return string(src^key)
}

func XORDecrypt(src, key int) string{

	return string(src^key)
}