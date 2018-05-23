package controllers

import (
	_ "bytes"
	_ "encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"time"
)

type WxHelperController struct {
	BaseController
}

func (this *WxHelperController) Get() {
	this.TplName = "product_edit.html"
	//this.Ctx.Output.Body([]byte(`你好，欢迎使用微信助手控制器`))
}

func (this *WxHelperController) Post() {

	act := this.Input().Get("act")
	//检查请求的方法
	if act != "" {
		switch act {
		//获取openid
		case "getOpenId":
			this.getOpenId()
			//发起支付
		// case "goPay":
		// 	this.goPay()

		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}
	}
	// this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
	// this.ServeJSON()

}

//获取openid--------------------------
func (this *WxHelperController) getOpenId() {

	appid := this.GetString("appid")
	secret := this.GetString("secret")
	js_code := this.GetString("js_code")
	grant_type := this.GetString("grant_type")
	// resp, err := http.PostForm("https://api.weixin.qq.com/sns/jscode2session", url.Values{"appid": {appid},
	// 	"secret":     {secret},
	// 	"js_code":    {js_code},
	// 	"grant_type": {grant_type}})
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=%v", appid, secret, js_code, grant_type))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(body)
	beego.Info(fmt.Sprintf(":::::::::::::%v", string(body)))
	// var result []byte
	// var ba base64.Encoding
	// ba.Decode(result, []byte(body))
	this.Data["json"] = map[string]interface{}{"status": 200, "openid": string(body), "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//发起支付---------------------------------------------------------
// func (this *WxHelperController) goPay() {
// 	var modOrder confirmOrder
// 	modOrder, ok := this.GetSession("confirmOrder").(confirmOrder)
// 	if !ok {
// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "确定订单获取失败", "time": time.Now().Format("2006-01-02 15:04:05")}
// 		this.ServeJSON()
// 		return
// 	}

// 	payType := this.GetString("PayType")
// 	if payType == "" {
// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少付款方式", "time": time.Now().Format("2006-01-02 15:04:05")}
// 		this.ServeJSON()
// 		return
// 	}

// 	//生成订单
// 	modOrder.TmpOrder.PayType = payType
// 	orderId, err := models.AddOrder(&modOrder.TmpOrder)
// 	if err != nil {
// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
// 		this.ServeJSON()
// 		return
// 	}
// 	//生成订单商品
// 	for _, item := range modOrder.ModProducts {
// 		var orderItem = new(models.TOrderItem)
// 		orderItem.OrderId = orderId
// 		orderItem.ProductId = item.Id
// 		orderItem.SumNum = item.BuyNum
// 		orderItem.SumPrice = item.SumPrice
// 		orderItem.CreateTime = time.Now().Format("2006-01-02 15:04:05")
// 		orderItem.SortId = item.SortId
// 		orderItem.Name = item.Name
// 		_, err = models.AddOrderItem(orderItem)
// 		if err != nil {
// 			models.DelOrder(orderId)
// 			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "订单项生成失败：" + err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
// 			this.ServeJSON()
// 			return
// 		}
// 	}
// 	this.DelSession("confirmOrder")
// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "添加订单成功", "time": time.Now().Format("2006-01-02 15:04:05")}
// 	this.ServeJSON()
// 	return

// }
