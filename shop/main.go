package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"shop/models"
	_ "shop/routers"
)

func init() {
	models.RegiesterDB()
}

func main() {
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//创建附件目录
	os.Mkdir("pictures", os.ModePerm)
	beego.SetStaticPath("pictures", "pictures")
	beego.Run()
}
