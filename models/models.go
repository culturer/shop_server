package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegiesterDB() {

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//数据库链接
	//注册默认数据库
	var db_url string = beego.AppConfig.String("username_DB") + ":" + beego.AppConfig.String("password_DB") + "@tcp(" + beego.AppConfig.String("host_DB") + ")/" + beego.AppConfig.String("name_DB") + "?charset=" + beego.AppConfig.String("charset")
	beego.Info(db_url)
	orm.RegisterDataBase("default", "mysql", db_url)
	// orm.RegisterDataBase("default", "mysql", "root:78901214@tcp(127.0.0.1:3306)/yoo_home?charset=utf8")
	//注册model
	orm.RegisterModel(new(TUser))
	orm.RegisterModel(new(TProductType))
	orm.RegisterModel(new(TProduct))
	orm.RegisterModel(new(TPicture))
	orm.RegisterModel(new(TPartner))
	orm.RegisterModel(new(TOrderItem))
	orm.RegisterModel(new(TOrder))
	orm.RegisterModel(new(TAddress))
	orm.RegisterModel(new(TAdvertise))
	orm.RegisterModel(new(TComment))

}
