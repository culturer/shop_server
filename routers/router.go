package routers

import (
	// "fmt"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"shop/controllers"
	_ "shop/models"
)

func init() {

	//初始化路由
	initRouter()
	//初始化过滤器
	// initFilter()

}

//初始化路由
func initRouter() {
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
	//分销商登录接口
	beego.Router("/p_login", &controllers.PLoginController{})
	//分销商采购接口
	beego.Router("/procurement", &controllers.ProcurementController{})
	//获取信息接口
	beego.Router("/wxhelper", &controllers.WxHelperController{})
	//广告接口
	beego.Router("/advertise", &controllers.AdvertiseController{})
}

//初始化过滤器
func initFilter() {
	// beego.InsertFilter("/*", beego.BeforeRouter, user_filter)
	//登录过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, login_filter)
	//供应商采购页面过滤器
	beego.InsertFilter("/procurement", beego.BeforeRouter, p_login_filter)
}

// func user_filter(ctx *context.Context) {
// 	//过滤的url表
// 	fileter_url := []string{"/login", "/register", "/products", "/get", "/wxhelper"}
// 	//pid --- 分销商Id
// 	for i := 0; i < len(fileter_url); i++ {
// 		if ctx.Request.RequestURI == fileter_url[i] {
// 			return
// 		}
// 	}
// 	_, ok := ctx.Input.Session("uid").(int64)
// 	if !ok {
// 		//beego.Info(fmt.Sprintf("redirect,uid:%v", uid))
// 		//ctx.Redirect(302, "/login")
// 		ctx.Output.Body([]byte(`{"status":"302","msg":"请重新登陆"}`))
// 	}
// }

//登录过滤器
func login_filter(ctx *context.Context) {

	//不过滤的url表
	n_fileter_url := []string{"/partner", "/login", "/register", "/products", "/advertise", "/get", "/procurement", "/p_login", "/wxhelper"}

	for i := 0; i < len(n_fileter_url); i++ {
		if ctx.Request.RequestURI == n_fileter_url[i] {
			return
		}
	}
	//uid --- 用户id
	_, ok := ctx.Input.Session("uid").(int64)
	if !ok {
		ctx.Output.Body([]byte(`{"status":"302","msg":"请重新登陆"}`))
	}

}

//gi
func p_login_filter(ctx *context.Context) {

	//过滤的url表
	fileter_url := []string{"/procurement"}
	//pid --- 分销商Id
	for i := 0; i < len(fileter_url); i++ {
		if ctx.Request.RequestURI == fileter_url[i] {
			_, ok := ctx.Input.Session("pid").(int64)
			if !ok {
				beego.Info("partnerId is null")
				//跳转到分销商登录页面
				ctx.Redirect(302, "/p_login")
			}
		}
	}

}
