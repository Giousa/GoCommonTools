/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/13
 */
package controllers

import "github.com/astaxie/beego"

type ToolsController struct {
	beego.Controller
}

func (this *ToolsController) BuildJavaMVCTemplate() {
	//packageName := this.GetString("packageName")
	//classNames := this.GetString("classNames")
	//author := this.GetString("author")
	//email := this.GetString("email")
	//returnBody := this.GetString("returnBody")
}

func (this *ToolsController) MysqlToStruct() {

}
