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
