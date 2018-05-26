package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"net/url"
	"shop/models"
	"strconv"
	"time"
)

type OrderController struct {
	BaseController
}

func (this *OrderController) Get() {
	var page string
	this.Ctx.Input.Bind(&page, "page")
	if page == "product_add" {
		this.TplName = "product_add.html"
	} else if page == "order_list" {
		this.TplName = "order_list.html"
	} else if page == "product_type_list" {
		this.TplName = "product_type_list.html"
	} else if page == "product_edit" {
		this.TplName = "product_edit.html"
	}
}

func (this *OrderController) Post() {

	act := this.Input().Get("act")
	//检查请求的方法
	if act != "" {
		switch act {
		//确认订单
		case "confirmOrder":
			this.confirmOrder()
			//生成订单
		case "createOrder":
			this.createOrder()
			//获取某个用户订单
		case "getOrderPageByUser":
			this.getOrderPageByUser()
			//分页获取订单
		case "getOrderPage":
			this.getOrderPage()
			//sql分页获取订单
		case "getOrderPageBySql":
			beego.Info("getOrderByUser")
			//修改订单
		case "editOrder":
			beego.Info("getOrderByUser")
			//退单
		case "cancelOrder":
			this.cancelOrder()
			//评论
		case "commentOrder":
			this.commentOrder()
			//发货
		case "goDlivery":
			this.goDlivery()
			//编辑物流状态
		case "editTranslateStatus":
			this.editTranslateStatus()
			//确定签收
		case "confirmSign":
			this.confirmSign()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}
	}

}

//生成session缓存确认订单 session key：confirmOrder--------------------------
func (this *OrderController) confirmOrder() {

	orderType, _ := strconv.Atoi(this.Input().Get("orderType"))

	products := this.GetString("products")
	if products == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少products参数", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	orderParam := this.GetString("orderParam")
	if orderParam == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少orderParam参数", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	var tmpProducts []models.TProduct
	var tmpOrder models.TOrder
	err := json.Unmarshal([]byte(products), &tmpProducts)
	err = json.Unmarshal([]byte(orderParam), &tmpOrder)
	beego.Info(tmpProducts)
	beego.Info(tmpOrder)
	if err != nil {

		beego.Info(tmpProducts)
		beego.Info(tmpOrder)

		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//后台核算商品------------------------------------------------------
	var modProducts []models.TProduct
	//订单总金额
	var payMoney float64
	for _, item := range tmpProducts {
		mod, err := models.GetProductById(item.Id)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		//数据处理
		mod.BuyNum = item.BuyNum
		if orderType == 1 {
			mod.SumPrice = float64(mod.BuyNum) * mod.StandardPrice
		} else {
			mod.SumPrice = float64(mod.BuyNum) * mod.Price
		}

		//添加到已处理的商品中
		modProducts = append(modProducts, *mod)
		payMoney += mod.SumPrice

	}
	//订单处理--------------------------------------------------------
	tmpOrder.ShouldPrice = payMoney
	tmpOrder.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	var sessionOrder = confirmOrder{ModProducts: modProducts, TmpOrder: tmpOrder}
	this.SetSession("confirmOrder", sessionOrder)

	this.Data["json"] = map[string]interface{}{"status": 200, "confirmOrder": sessionOrder, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//生成用户订单---------------------------------------------------------
func (this *OrderController) createOrder() {
	var modOrder confirmOrder
	modOrder, ok := this.GetSession("confirmOrder").(confirmOrder)
	if !ok {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "确定订单获取失败", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	payType := this.GetString("PayType")
	if payType == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少付款方式", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//生成订单
	modOrder.TmpOrder.PayType = payType
	orderId, err := models.AddOrder(&modOrder.TmpOrder)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//生成订单商品
	for _, item := range modOrder.ModProducts {
		var orderItem = new(models.TOrderItem)
		orderItem.OrderId = orderId
		orderItem.ProductId = item.Id
		orderItem.SumNum = item.BuyNum
		orderItem.SumPrice = item.SumPrice
		orderItem.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		orderItem.SortId = item.SortId
		orderItem.Name = item.Name
		_, err = models.AddOrderItem(orderItem)
		if err != nil {
			models.DelOrder(orderId)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "订单项生成失败：" + err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}
	this.DelSession("confirmOrder")
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "添加订单成功", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//分页获取订单列表------------------------------------------------------
func (this *OrderController) getOrderPage() {
	where := this.GetString("where")
	size, _ := strconv.Atoi(this.GetString("size"))
	index, _ := strconv.Atoi(this.GetString("index"))
	if size == 0 || index == 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "getOrderPage参数不齐", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单列表
	orderList, count, err := models.GetOrderPage(index, size, where)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//构造返回订单
	var completeOrderList []completeOrder
	for _, item := range orderList {
		orderIntems, itemCount, _err := models.GetOrderItemPage(1, 0, fmt.Sprintf(" and order_id=%v", item.Id))
		if _err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": _err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		var modCompleteOrder = completeOrder{OrderInfo: item, ItemCount: itemCount, OrderItems: orderIntems}
		completeOrderList = append(completeOrderList, modCompleteOrder)

	}
	this.Data["json"] = map[string]interface{}{"status": 200, "count": count, "dataList": completeOrderList, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//分页用户订单列表------------------------------------------------------
func (this *OrderController) getOrderPageByUser() {
	where := this.GetString("where")
	size, _ := strconv.Atoi(this.GetString("size"))
	index, _ := strconv.Atoi(this.GetString("index"))
	if size == 0 || index == 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "getOrderPage参数不齐", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单列表
	orderList, count, err := models.GetOrderPage(index, size, fmt.Sprintf(" and user_id=%v %v", this.GetSession("uid"), where))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//构造返回订单
	var completeOrderList []completeOrder
	for _, item := range orderList {
		orderIntems, itemCount, _err := models.GetOrderItemPage(1, 0, fmt.Sprintf(" and order_id=%v", item.Id))
		if _err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": _err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		var modCompleteOrder = completeOrder{OrderInfo: item, ItemCount: itemCount, OrderItems: orderIntems}
		completeOrderList = append(completeOrderList, modCompleteOrder)

	}
	this.Data["json"] = map[string]interface{}{"status": 200, "count": count, "dataList": completeOrderList, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//退单-----------------------------------------------------------
func (this *OrderController) cancelOrder() {
	cancelComments := this.GetString("comments")
	orderId, _ := strconv.Atoi(this.GetString("orderId"))

	if cancelComments == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少退单理由", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单
	order, err := models.GetOrderById(int64(orderId))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "找不到订单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//修改订单为退款状态
	order.IsCancel = true
	order.CancelComments = url.QueryEscape(fmt.Sprintf("时间:%v \n %v", time.Now().Format("2006-01-02 15:04:05"), cancelComments))
	models.EditOrder(order)
	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//评论订单-----------------------------------------------------------
func (this *OrderController) commentOrder() {
	comments := this.GetString("comments")
	orderId, _ := strconv.Atoi(this.GetString("orderId"))

	if comments == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少评论", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单
	order, err := models.GetOrderById(int64(orderId))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "找不到订单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//修改订单为评论状态
	order.IsComment = true
	order.Comments = url.QueryEscape(fmt.Sprintf("时间:%v \n <br/> %v \n <br/>", time.Now().Format("2006-01-02 15:04:05"), comments))
	models.EditOrder(order)
	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//发货
func (this *OrderController) goDlivery() {

	comments := this.GetString("comments")
	orderId, _ := strconv.Atoi(this.GetString("orderId"))

	if comments == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "填写发货单号", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单
	order, err := models.GetOrderById(int64(orderId))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "找不到订单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//修改订单为发货状态
	order.IsDlivery = true
	order.TranslateStatus = url.QueryEscape(fmt.Sprintf("时间:%v \n <br/> %v \n <br/>", time.Now().Format("2006-01-02 15:04:05"), comments))
	models.EditOrder(order)
	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//编辑物流状态
func (this *OrderController) editTranslateStatus() {

	comments := this.GetString("comments")
	orderId, _ := strconv.Atoi(this.GetString("orderId"))

	if comments == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "填写物流信息", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取订单
	order, err := models.GetOrderById(int64(orderId))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "找不到订单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//修改物流状态

	order.TranslateStatus += url.QueryEscape(fmt.Sprintf("时间:%v \n <br/> %v \n <br/>", time.Now().Format("2006-01-02 15:04:05"), comments))
	models.EditOrder(order)
	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//确定签收
func (this *OrderController) confirmSign() {

	comments := this.GetString("comments")
	orderId, _ := strconv.Atoi(this.GetString("orderId"))

	//获取订单
	order, err := models.GetOrderById(int64(orderId))
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "找不到订单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//修改订单状态
	order.IsSign = true
	if comments != "" {
		order.IsComment = true
		order.Comments = url.QueryEscape(fmt.Sprintf("时间:%v \n <br/> %v \n <br/>", time.Now().Format("2006-01-02 15:04:05"), comments))

	}
	models.EditOrder(order)
	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//确定订单结构
type confirmOrder struct {
	ModProducts []models.TProduct
	TmpOrder    models.TOrder
}

//返回订单列表结构
type completeOrder struct {
	OrderInfo  *models.TOrder
	OrderItems []*models.TOrderItem
	ItemCount  int
}

// func (this *OrderController) addOrder(userId int64, addressId int64, address string, payType string, realPrice float64, shouldPrice float64, priceMsg string, partnerId int64) (int64, error) {

// 	orderId, err := models.AddOrder(userId, addressId, address, payType, realPrice, shouldPrice, priceMsg, partnerId)
// 	return orderId, err
// }

// func (this *OrderController) delOrder(orderId int64) error {
// 	err := models.DelOrder(orderId)
// 	return err
// }

// func (this *OrderController) getOrderById(orderId int64) (*models.TOrder, error) {
// 	order, err := models.GetOrderById(orderId)
// 	return order, err
// }

// func (this *OrderController) getOrderByUserId(userId int64, pageNo, pageSize int, where string) ([]*models.TOrder, int, error) {
// 	orders, totalPage, err := models.GetOrderByUserId(userId, pageNo, pageSize, where)
// 	return orders, totalPage, err
// }

// func (this *OrderController) getOrderByUserIdS(userId int64, orderStatus int, pageNo, pageSize int) ([]*models.TOrder, int, error) {
// 	orders, totalPage, err := models.GetOrderByUserIdS(userId, orderStatus, pageNo, pageSize)
// 	return orders, totalPage, err
// }

// func (this *OrderController) getOrderByPartnerId(partnerId int64, pageNo, pageSize int, where string) ([]*models.TOrder, int, error) {
// 	orders, totalPage, err := models.GetOrderByPartnerId(partnerId, pageNo, pageSize, where)
// 	return orders, totalPage, err
// }

// func (this *OrderController) getOrderByPartnerIdS(partnerId int64, orderStatus int, pageNo, pageSize int) ([]*models.TOrder, int, error) {
// 	orders, totalPage, err := models.GetOrderByIdPartnerIdS(partnerId, pageNo, pageSize, orderStatus)
// 	return orders, totalPage, err
// }

// func (this *OrderController) mdfyOrderStatus(orderId int64, orderStatus int, payType string) error {
// 	err := models.MdfyOrderStatus(orderId, orderStatus, payType)
// 	return err
// }

// func (this *OrderController) mdfyOrderAddress(orderId int64, addressId int64) error {
// 	err := models.MdfyOrderAddress(orderId, addressId)
// 	return err
// }

// func (this *OrderController) mdfyOrderSort(orderId int64, sortId int) error {
// 	err := models.MdfyOrderSort(orderId, sortId)
// 	return err
// }

// func (this *OrderController) mdfyOrderTranslateStatus(orderId int64, status string) error {
// 	err := models.MdfyOrderTranslateStatus(orderId, status)
// 	return err
// }

// func (this *OrderController) getOrderItemById(orderItemId int64) (*models.TOrder, error) {
// 	// orderItem, err := models.GetOrderItemById(orderItemId)
// 	return nil, nil
// }

// func (this *OrderController) getOrderItemByProductId(productId int64, pageNo, pageSize int, where string) ([]*models.TOrderItem, int, error) {
// 	orderItems, totalPage, err := models.GetOrderItemByProductId(productId, pageNo, pageSize, "")
// 	return orderItems, totalPage, err
// }

// func (this *OrderController) getOrderItemByOrderId(orderId int64, pageNo, pageSize int, where string) ([]*models.TOrderItem, int, error) {
// 	orderItems, totalPage, err := models.GetOrderItemByOrderId(orderId, pageNo, pageSize, "")
// 	return orderItems, totalPage, err
// }

// func (this *OrderController) addOrderItem(productId int64, orderId int64) (int64, error) {
// 	orderItemId, err := models.AddOrderItem(productId, orderId)
// 	return orderItemId, err
// }

// func (this *OrderController) delOrderItem(orderItemId int64) error {
// 	err := models.DelOrderItem(orderItemId)
// 	return err
// }
