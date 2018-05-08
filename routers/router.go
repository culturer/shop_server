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
	//获取openId
	// beego.Router("/ope", &controllers.OpenIdController{})
	var FilterUser = func(ctx *context.Context) {

		if ctx.Request.RequestURI != "/login" {
			if ctx.Request.RequestURI != "/register" {
				if ctx.Request.RequestURI != "/products" {
					_, ok := ctx.Input.Session("uid").(int64)
					if !ok {
						//beego.Info(fmt.Sprintf("redirect,uid:%v", uid))
						//ctx.Redirect(302, "/login")
						ctx.Output.Body([]byte(`{"status":"302","msg":"请重新登陆"}`))
					}
				}

			}

		}

	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}
