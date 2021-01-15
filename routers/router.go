package routers

import (
	"common-tools/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //生成java MVC 模板
	beego.Router("/buildJavaMVCTemplate", &controllers.ToolsController{},"get:BuildJavaMVCTemplate")
	beego.Router("/downloadJavaMVCTemplate", &controllers.ToolsController{},"get:DownloadJavaMVCTemplate")

    //mysql导出struct
    beego.Router("/mysqlToStruct", &controllers.ToolsController{},"get:MysqlToStruct")


    //加密和解密
    beego.Router("/encryption", &controllers.ToolsController{},"get:EncryptionData")

	//Excel解析

	//上传图片

	//Go工具类方法

	//生成二维码
}
