/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/13
 */
package controllers

import (
	"common-tools/service"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ToolsController struct {
	beego.Controller
}

//http://localhost:8282/buildJavaMVCTemplate
func (this *ToolsController) BuildJavaMVCTemplate() {
	storageAddress := this.GetString("storageAddress")
	packageName := this.GetString("packageName")
	classNames := this.GetString("classNames")
	author := this.GetString("author")
	email := this.GetString("email")
	resultBody := this.GetString("resultBody")

	if storageAddress == "" || packageName == "" || classNames == ""{
		fmt.Println("路径、包名、类名不能为空")
		this.Ctx.WriteString("路径、包名、类名不能为空")
		return
	}

	javaModel := service.JavaModel{
		StorageAddress:storageAddress,
		Package: packageName,
		NameMap: strings.Split(classNames,","),
		ResultBody: resultBody,
		Author: author,
		Email: email,
	}

	fmt.Println(javaModel)

	javaModel.Run()

	this.Ctx.WriteString("ok")
}

//http://localhost:8282/mysqlToStruct
func (this *ToolsController) MysqlToStruct() {

	storageAddress := this.GetString("storageAddress")
	fileName := this.GetString("fileName")
	ip := this.GetString("ip")
	root := this.GetString("root")
	password := this.GetString("password")
	dbName := this.GetString("dbName")

	if  ip == "" || root == "" || password == "" || dbName == ""{
		fmt.Println("参数缺失")
		this.Ctx.WriteString("参数缺失")
		return
	}

	if storageAddress == ""{
		storageAddress = "."
	}


	if fileName == ""{
		fileName = "mysqlToStruct.go"
	}

	//mysqlModel := service.MysqlModel{
	//	IP: "47.103.115.252",
	//	Root: "root",
	//	Password: "h5s/X_7FLkzj",
	//	DbName: "das",
	//	StorageAddress: ".",
	//	FileName: "mysqlToStruct.go",
	//}
	mysqlModel := service.MysqlModel{
		IP: ip,
		Root: root,
		Password: password,
		DbName: dbName,
		StorageAddress: storageAddress,
		FileName: fileName,
	}


	result,_ := mysqlModel.InitMysql()
	fmt.Println("返回结果：",result)
	this.Ctx.WriteString(result)
}

//加密解密
//http://localhost:8282/encryption
func (this *ToolsController) EncryptionData()  {

	verify,_ := this.GetBool("verify",false)
	isDecrypt,_ := this.GetBool("isDecrypt",false)
	encryptKey := this.GetString("encryptKey")
	encryptSource := this.GetString("encryptSource")
	encryptType := this.GetString("encryptType")
	encryptResult := this.GetString("encryptResult")

	fmt.Println("是否校验:",verify)
	fmt.Println("Key:",encryptKey)
	fmt.Println("原数据:",encryptSource)
	fmt.Println("方法:",encryptType)
	fmt.Println("是否解密:",isDecrypt)
	fmt.Println("加密后的字符串:",encryptResult)

	if  encryptType == ""{
		fmt.Println("参数缺失")
		this.Ctx.WriteString("参数缺失")
		return
	}

	var result string
	switch encryptType {
		case "md5":
			result = service.MD5("今天is好day")
			break

		case "hmac":
			key := []byte(encryptKey)
			src := []byte(encryptSource)

			if verify{
				//校验
				flag := service.VerifyHMAC([]byte(encryptResult), src, key)
				if flag{
					result = "ok"
				}else{
					result = "wrong"
				}

			}else{
				//加密

				result = service.HMAC(src, key)
			}

			break
		case "sha":
			res,_ := service.Sha1("/Users/zhangmengmeng/Downloads/temp/mysqlToStruct.go")
			result = res
			break
		case "3des":
			if isDecrypt{
				//解密
				result = service.TripleDESDecrypt([]byte(encryptResult),[]byte(encryptKey))
			}else{
				//加密
				result = service.TripleDESEncrypt([]byte(encryptSource),[]byte(encryptKey))
			}
			break
		case "ase":
			if isDecrypt{
				//解密
				result = service.AESDecrypt([]byte(encryptResult),[]byte(encryptKey))
			}else{
				//加密
				result = service.AESEncrypt([]byte(encryptSource),[]byte(encryptKey))
			}
			break
		case "des":
			if isDecrypt{
				//解密
				result = service.DesDecrypt([]byte(encryptResult),[]byte(encryptKey))
			}else{
				//加密
				result = service.DesEncrypt([]byte(encryptSource),[]byte(encryptKey))
			}
			break
		case "xor":
			src,_ := strconv.Atoi(encryptSource)
			res,_ := strconv.Atoi(encryptResult)
			key,_ := strconv.Atoi(encryptKey)

			if isDecrypt{
				//解密
				result = service.XORDecrypt(res,key)
			}else{
				//加密
				result = service.XOREncrypt(src,key)
			}
			break
		case "rsasign":

			break
		case "rsa":

			break
	}
	fmt.Println(result)
	this.Ctx.WriteString(result)
}

//http://localhost:8282/downloadBuildFile
//浏览器请求成功后，会以downloadJavaMVCTemplate.zip名称下载到本地
func (this *ToolsController) DownloadBuildFile() {

	fileName := this.GetString("fileName")

	fmt.Println("准备下载......")

	//filePath := "/Users/zhangmengmeng/Downloads/temp/com.zip"
	var filePath string
	if fileName == "files.zip"{
		filePath = "/usr/local/workspace/"+fileName
	}else{
		//hello.go
		filePath = "/usr/local/workspace/structFiles/"+fileName
	}
	fmt.Println("文件路径：",filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件不存在")
		this.Ctx.WriteString("文件不存在")
		return
	}
	defer f.Close()

	// 将文件读取出来
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("文件读取失败")
		this.Ctx.WriteString("文件读取失败")
		return
	}

	//fmt.Println("--------------文件内容--------------")
	//fmt.Println(string(data))
	fmt.Println("文件下载中...")
	this.Ctx.WriteString(string(data))

	fmt.Println("下载完毕")

	//删除zip
	os.Remove(filePath)

}
