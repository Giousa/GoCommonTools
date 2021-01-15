/**
 *@Desc: 加密文件
 *@Author:Giousa
 *@Date:2021/1/14
 */
package encry

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

//使用sha1计算文件指纹
func GetSha1(filePath string) (string,error) {
	// 1. 打开文件
	fmt.Println("加密文件地址：",filePath)
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败,err = ",err)
		return "文件打开失败",err
	}
	// 2. 创建基于sha1算法的Hash对象
	myHash := sha1.New()
	// 3. 将文件数据拷贝给哈希对象
	num, err := io.Copy(myHash, fp)
	if err != nil {
		fmt.Println("拷贝文件失败,err = ",err)
		return "拷贝文件失败",err
	}
	fmt.Println("文件大小: ", num)
	// 4. 计算文件的哈希值
	tmp1 := myHash.Sum(nil)
	// 5. 数据格式转换
	result := hex.EncodeToString(tmp1)
	fmt.Println("sha1: ", result)

	return result,nil
}
