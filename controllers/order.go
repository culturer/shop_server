package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"shop/models"
	_ "strconv"
	"time"
)

type OrderController struct {
	BaseController
}

func (this *OrderController) Get() {
	// this.Data["Website"] = "beego.me"
	// this.Data["Email"] = "astaxie@gmail.com"

	this.TplName = "order.html"
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
			beego.Info("createOrder")
			//获取某个用户订单
		case "getOrderByUser":
			beego.Info("getOrderByUser")
			//分页获取订单
		case "getOrderPage":
			beego.Info("getOrderByUser")
			//sql分页获取订单
		case "getOrderPageBySql":
			beego.Info("getOrderByUser")
			//修改订单
		case "editOrder":
			beego.Info("getOrderByUser")
			//退单
		case "cancelOrder":
			beego.Info("getOrderByUser")
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()

		}
	}

}
func (this *OrderController) confirmOrder() {

	products := this.GetString("products")
	if products == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少products参数", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	orderParam := this.GetString("orderParam")
	if orderParam == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "缺少orderParam参数", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
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
		}
		//数据处理
		mod.BuyNum = item.BuyNum
		mod.SumPrice = float64(mod.BuyNum) * mod.Price
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

}

type confirmOrder struct {
	ModProducts []models.TProduct
	TmpOrder    models.TOrder
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
