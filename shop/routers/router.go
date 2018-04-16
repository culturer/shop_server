package routers

import (
	"github.com/astaxie/beego"
	"shop/controllers"
)

func init() {
	//登录
	beego.Router("/login", &controllers.LoginController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	//商品相关
	beego.Router("/products", &controllers.ProductController{})
	//订单相关
	beego.Router("/orders", &controllers.OrderController{})
}
