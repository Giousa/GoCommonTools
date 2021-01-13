package routers

import (
	"common-tools/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //生成java MVC 模板
	beego.Router("/buildJavaMVCTemplate", &controllers.ToolsController{},"get:BuildJavaMVCTemplate")

    //mysql导出struct
    beego.Router("/mysqlToStruct", &controllers.ToolsController{},"get:MysqlToStruct")

	//Excel解析

	//上传图片

	//Go工具类方法

	//生成二维码
}
