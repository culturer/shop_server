package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"
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
	// [types == 0  获取订单相关]
	// [types == 1  获取订单商品项相关]
	types, _ := strconv.Atoi(this.Input().Get("types"))
	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	options, _ := strconv.Atoi(this.Input().Get("options"))

	if types == 0 {

		if options == 0 {
			// [getType == 0  根据orderId查询]
			// [getType == 1  根据userId查询]
			// [getType == 2  根据orderId,orderStatus查询]
			// [getType == 3  根据partnerId查询]
			// [getType == 4  根据partnerId,orderStatus查询]
			getType, _ := strconv.Atoi(this.Input().Get("getType"))

			if getType == 0 {
				orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
				order, err := this.getOrderById(orderId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 1 {
				userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
				orders, err := this.getOrderByUserId(userId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 2 {
				userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
				orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
				orders, err := this.getOrderByUserIdS(userId, orderStatus)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 3 {
				partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
				orders, err := this.getOrderByPartnerId(partnerId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 4 {
				partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
				orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
				orders, err := this.getOrderByPartnerIdS(partnerId, orderStatus)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	if options == 1 {
		userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
		orderId, err := this.addOrder(userId, addressId, partnerId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 添加订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "orderId": orderId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
		err := this.delOrder(orderId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 删除订单成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		// [mdfyType == 0  修改订单状态]
		// [mdfyType == 1  修改订单地址]
		mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
		if mdfyType == 0 {
			orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			err := this.mdfyOrderStatus(orderId, orderStatus)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改订单成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		if mdfyType == 1 {
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
			err := this.mdfyOrderAddress(orderId, addressId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改订单成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}

	if types == 1 {

		if options == 0 {
			// [getType == 0  根据orderItemId查询]
			// [getType == 1  根据productId查询]
			// [getType == 2  根据orderId查询]
			getType, _ := strconv.Atoi(this.Input().Get("getType"))

			if getType == 0 {
				orderItemId, _ := strconv.ParseInt(this.Input().Get("orderItemId"), 10, 64)
				orderItem, err := this.getOrderItemById(orderItemId)
				if err != nil {
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orderItem": orderItem, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 1 {
				productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
				orderItems, err := this.getOrderItemByProductId(productId)
				if err != nil {
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orderItems": orderItems, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 2 {
				orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
				orderItems, err := this.getOrderItemByOrderId(orderId)
				if err != nil {
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orderItems": orderItems, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		}

		if options == 1 {
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)

			orderItemId, err := this.addOrderItem(productId, orderId)

			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "增加订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "orderItemId": orderItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if options == 2 {
			orderItemId, _ := strconv.ParseInt(this.Input().Get("orderItemId"), 10, 64)

			err := this.delOrderItem(orderItemId)

			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "删除订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "删除订单项成功", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if options == 3 {
			beego.Info("暂不提供修改接口")
		}

	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 参数异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *OrderController) addOrder(userId int64, addressId int64, partnerId int64) (int64, error) {
	orderId, err := models.AddOrder(userId, addressId, partnerId)
	return orderId, err
}

func (this *OrderController) delOrder(orderId int64) error {
	err := models.DelOrder(orderId)
	return err
}

func (this *OrderController) getOrderById(orderId int64) (*models.TOrder, error) {
	order, err := models.GetOrderById(orderId)
	return order, err
}

func (this *OrderController) getOrderByUserId(userId int64) ([]*models.TOrder, error) {
	orders, err := models.GetOrderByUserId(userId)
	return orders, err
}

func (this *OrderController) getOrderByUserIdS(userId int64, orderStatus int) ([]*models.TOrder, error) {
	orders, err := models.GetOrderByUserIdS(userId, orderStatus)
	return orders, err
}

func (this *OrderController) getOrderByPartnerId(partnerId int64) ([]*models.TOrder, error) {
	orders, err := models.GetOrderByPartnerId(partnerId)
	return orders, err
}

func (this *OrderController) getOrderByPartnerIdS(partnerId int64, orderStatus int) ([]*models.TOrder, error) {
	orders, err := models.GetOrderByIdPartnerIdS(partnerId, orderStatus)
	return orders, err
}

func (this *OrderController) mdfyOrderStatus(orderId int64, orderStatus int) error {
	err := models.MdfyOrderStatus(orderId, orderStatus)
	return err
}

func (this *OrderController) mdfyOrderAddress(orderId int64, addressId int64) error {
	err := models.MdfyOrderAddress(orderId, addressId)
	return err
}

func (this *OrderController) getOrderItemById(orderItemId int64) (*models.TOrder, error) {
	orderItem, err := models.GetOrderItemById(orderItemId)
	return orderItem, err
}

func (this *OrderController) getOrderItemByProductId(productId int64) ([]*models.TOrderItem, error) {
	orderItems, err := models.GetOrderItemByProductId(productId)
	return orderItems, err
}

func (this *OrderController) getOrderItemByOrderId(orderId int64) ([]*models.TOrderItem, error) {
	orderItems, err := models.GetOrderItemByOrderId(orderId)
	return orderItems, err
}

func (this *OrderController) addOrderItem(productId int64, orderId int64) (int64, error) {
	orderItemId, err := models.AddOrderItem(productId, orderId)
	return orderItemId, err
}

func (this *OrderController) delOrderItem(orderItemId int64) error {
	err := models.DelOrderItem(orderItemId)
	return err
}
