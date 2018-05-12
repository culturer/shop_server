package controllers

import (
	"github.com/astaxie/beego"
	"shop/bean"
	"strconv"
	"time"
)

type GetController struct {
	BaseController
}

func (this *GetController) Get() {
	// this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "get_test.html"
}

func (this *GetController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "*")
	// [options == 0  获取订单]
	// [options == 1  获取订单项]
	// [options == 2  获取分销商列表]
	// [options == 3  获取产品]
	// [options == 4  获取产品分类]
	// [options == 5  获取用户信息]
	options, _ := strconv.Atoi(this.Input().Get("options"))
	if options == 0 {

		// orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
		// order, err := bean.GetOrderBean(orderId)
		// if err != nil {
		// 	beego.Info(err.Error())
		// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取订单信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		// 	this.ServeJSON()
		// 	return
		// }
		// this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
		// this.ServeJSON()
		// return

	}

	if options == 1 {

		// orderItemId, _ := strconv.ParseInt(this.Input().Get("orderItemId"), 10, 64)
		// orderItem, err := bean.GetOrderItemBean(orderItemId)
		// if err != nil {
		// 	beego.Info(err.Error())
		// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取订单项信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		// 	this.ServeJSON()
		// 	return
		// }
		// this.Data["json"] = map[string]interface{}{"status": 200, "orderItem": orderItem, "time": time.Now().Format("2006-01-02 15:04:05")}
		// this.ServeJSON()
		// return

	}

	if options == 2 {

		// partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
		// partner, err := bean.GetPartnerBean(partnerId)
		// if err != nil {
		// 	beego.Info(err.Error())
		// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取分销商信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		// 	this.ServeJSON()
		// 	return
		// }
		// this.Data["json"] = map[string]interface{}{"status": 200, "partner": partner, "time": time.Now().Format("2006-01-02 15:04:05")}
		// this.ServeJSON()
		// return

	}

	if options == 3 {

		productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
		product, err := bean.GetProductBean(productId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "product": product, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	if options == 4 {

		productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
		productType, err := bean.GetProductTypeBean(productTypeId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品分类信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "productType": productType, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	if options == 5 {

		// userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
		// user, err := bean.GetUserBean(userId)
		// if err != nil {
		// 	beego.Info(err.Error())
		// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取用户信息失败，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		// 	this.ServeJSON()
		// 	return
		// }
		// this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
		// this.ServeJSON()
		// return

	}

}
