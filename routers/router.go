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
	beego.Router("/order", &controllers.OrderController{})
	//分销商相关
	beego.Router("/partner", &controllers.PartnerController{})
	//图片相关
	beego.Router("/picture", &controllers.PictureController{})
	//用户相关
	beego.Router("/user", &controllers.UserController{})
	//地址相关
	beego.Router("/address", &controllers.AddressController{})
	//获取信息接口
	beego.Router("/get", &controllers.GetController{})
}
