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
		this.Ctx.WriteString("包名或类名不能为空")
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

	//javaModel.Run()

	this.Ctx.WriteString("成功")
}

func (this *ToolsController) MysqlToStruct() {

}
