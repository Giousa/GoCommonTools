/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/13
 */
package service

import (
	templates "common-tools/models"
	"common-tools/utils"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

type JavaModel struct {
	StorageAddress string //文件存储地址
	Package string //包名
	NameMap []string //类名
	ResultBody string//返回值
	Author string //作者
	Email string //邮箱
}

type ClassTemplate struct {
	Package string
	Name string
	NameHumpLower string //驼峰，首字母小写
	NameSeparate string //蛇形，-分隔
	ResultBody string//返回值
	Author string //作者
	Email string //邮箱
	DateTime string //时间
}

func (j *JavaModel) Run() {
	packageName := j.Package
	fmt.Println("包名：",packageName)
	//path,_ := os.Getwd()
	//path := j.StorageAddress
	path := "/usr/local/workspace/files"
	fmt.Println("存储目录：",path)

	//为防止冲突，删除之前文件
	os.RemoveAll(path)
	os.MkdirAll(path,os.ModePerm)

	//os.Mkdir(packageName,os.FileMode())
	newPackage := strings.Replace(packageName,".","/",-1)
	fmt.Println("包名转文件目录："+newPackage)
	fmt.Println("开始在根目录下创建包名对应文件夹：")
	newPath := path+"/"+newPackage



	os.MkdirAll(newPath,os.ModePerm)

	fmt.Println("开始创建：controller service service/impl文件夹")
	os.MkdirAll(newPath+"/controller",os.ModePerm)
	os.MkdirAll(newPath+"/service/impl",os.ModePerm)

	pathController := newPath+"/controller/"
	pathService := newPath+"/service/"
	pathImpl := newPath+"/service/impl/"

	fmt.Println(pathController)
	fmt.Println(pathService)
	fmt.Println(pathImpl)
	for _,v := range j.NameMap{
		if j.ResultBody == ""{
			j.ResultBody = "ResultVO"
		}
		buildControllerFile(pathController,packageName,j.ResultBody,j.Author,j.Email,v)
		buildServiceFile(pathService,packageName,j.ResultBody,j.Author,j.Email,v)
		buildServiceImplFile(pathImpl,packageName,j.ResultBody,j.Author,j.Email,v)
	}


	//将path文件夹下新生成的目录文件，压缩
	//nameList := strings.Split(packageName,".")
	//if nameList != nil && len(nameList) >0 {
	//	///usr/local/workspace/files/com
	//	//zipPath := path+"/"+nameList[0]
	//	zipPath := "/usr/local/workspace/files/"+nameList[0]
	//	//开始压缩
	//	utils.ZipDir(zipPath,"/usr/local/workspace/files/zip/"+nameList[0]+".zip")
	//
	//}

	utils.ZipDir("/usr/local/workspace/files/","/usr/local/workspace/files.zip")

}

func buildControllerFile(path string, packageName string, resultBody string,author string,email string,name string)  {
	className := name+"Controller.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Controller文件：",file.Name())

	t := template.Must(template.New("").Parse(templates.TextController))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: utils.Lcfirst(name),
		NameSeparate: utils.SeparateToString(name),
		ResultBody: resultBody,
		DateTime:time.Now().Format("2006-01-02"),
		Author: author,
		Email: email,
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

func buildServiceFile(path string, packageName string, resultBody string,author string,email string,name string)  {
	className := name+"Service.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Service文件：",file.Name())

	t := template.Must(template.New("").Parse(templates.TextService))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: utils.Lcfirst(name),
		NameSeparate: utils.SeparateToString(name),
		ResultBody: resultBody,
		DateTime:time.Now().Format("2006-01-02"),
		Author: author,
		Email: email,
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

func buildServiceImplFile(path string, packageName string, resultBody string,author string,email string,name string)  {
	className := name+"ServiceImpl.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Impl文件：",file.Name())

	t := template.Must(template.New("").Parse(templates.TextServiceImpl))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: utils.Lcfirst(name),
		NameSeparate: utils.SeparateToString(name),
		ResultBody: resultBody,
		DateTime:time.Now().Format("2006-01-02"),
		Author: author,
		Email: email,
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

