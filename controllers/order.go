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
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orders, totalPage, err := this.getOrderByUserId(userId, pageNo, pageSize, "")
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 2 {
				userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
				orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orders, totalPage, err := this.getOrderByUserIdS(userId, orderStatus, pageNo, pageSize)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 3 {
				partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orders, totalPage, err := this.getOrderByPartnerId(partnerId, pageNo, pageSize, "")
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 4 {
				partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
				orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orders, totalPage, err := this.getOrderByPartnerIdS(partnerId, orderStatus, pageNo, pageSize)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询订单异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orders": orders, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
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
		payType := this.Input().Get("payType")
		address := this.Input().Get("address")
		priceMsg := this.Input().Get("priceMsg")
		realPrice, _ := strconv.ParseFloat(this.Input().Get("realPrice"), 32)
		shouldPrice, _ := strconv.ParseFloat(this.Input().Get("shouldPrice"), 32)

		orderId, err := this.addOrder(userId, addressId, address, payType, realPrice, shouldPrice, priceMsg, partnerId)
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
		// [mdfyType == 2  修改订单优先级]
		// [mdfyType == 3  修改订单物流状态]
		mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
		if mdfyType == 0 {
			orderStatus, _ := strconv.Atoi(this.Input().Get("orderStatus"))
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			//isStatus 订单的具体状态
			isStatus, _ := strconv.ParseBool(this.Input().Get("isStatus"))
			//状态标志位
			isFlag, _ := strconv.Atoi(this.Input().Get("isFlag"))
			payType := this.Input().Get("payType")
			switch isFlag {
			case 0:
				models.OrderIsPay(orderId, isStatus)
			case 1:
				models.OrderIsDlivery(orderId, isStatus)
			case 2:
				models.OrderIsSign(orderId, isStatus)
			case 3:
				models.OrderIsCash(orderId, isStatus)
			case 4:
				models.OrderIsComment(orderId, isStatus)
			default:
			}
			err := this.mdfyOrderStatus(orderId, orderStatus, payType)
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
		if mdfyType == 2 {
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			sortId, _ := strconv.Atoi(this.Input().Get("sortId"))
			err := this.mdfyOrderSort(orderId, sortId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改订单优先级,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改订单优先级成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}
		if mdfyType == 3 {
			orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
			status := this.Input().Get("status")
			err := this.mdfyOrderTranslateStatus(orderId, status)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改订单物流状态,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改订单物流状态成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
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
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orderItems, totalPage, err := this.getOrderItemByProductId(productId, pageNo, pageSize, "")
				if err != nil {
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orderItems": orderItems, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 2 {
				orderId, _ := strconv.ParseInt(this.Input().Get("orderId"), 10, 64)
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				orderItems, totalPage, err := this.getOrderItemByOrderId(orderId, pageNo, pageSize, "")
				if err != nil {
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询订单项异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "orderItems": orderItems, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
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

func (this *OrderController) addOrder(userId int64, addressId int64, address string, payType string, realPrice float64, shouldPrice float64, priceMsg string, partnerId int64) (int64, error) {
	orderId, err := models.AddOrder(userId, addressId, address, payType, realPrice, shouldPrice, priceMsg, partnerId)
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

func (this *OrderController) getOrderByUserId(userId int64, pageNo, pageSize int, where string) ([]*models.TOrder, int, error) {
	orders, totalPage, err := models.GetOrderByUserId(userId, pageNo, pageSize, where)
	return orders, totalPage, err
}

func (this *OrderController) getOrderByUserIdS(userId int64, orderStatus int, pageNo, pageSize int) ([]*models.TOrder, int, error) {
	orders, totalPage, err := models.GetOrderByUserIdS(userId, orderStatus, pageNo, pageSize)
	return orders, totalPage, err
}

func (this *OrderController) getOrderByPartnerId(partnerId int64, pageNo, pageSize int, where string) ([]*models.TOrder, int, error) {
	orders, totalPage, err := models.GetOrderByPartnerId(partnerId, pageNo, pageSize, where)
	return orders, totalPage, err
}

func (this *OrderController) getOrderByPartnerIdS(partnerId int64, orderStatus int, pageNo, pageSize int) ([]*models.TOrder, int, error) {
	orders, totalPage, err := models.GetOrderByIdPartnerIdS(partnerId, pageNo, pageSize, orderStatus)
	return orders, totalPage, err
}

func (this *OrderController) mdfyOrderStatus(orderId int64, orderStatus int, payType string) error {
	err := models.MdfyOrderStatus(orderId, orderStatus, payType)
	return err
}

func (this *OrderController) mdfyOrderAddress(orderId int64, addressId int64) error {
	err := models.MdfyOrderAddress(orderId, addressId)
	return err
}

func (this *OrderController) mdfyOrderSort(orderId int64, sortId int) error {
	err := models.MdfyOrderSort(orderId, sortId)
	return err
}

func (this *OrderController) mdfyOrderTranslateStatus(orderId int64, status string) error {
	err := models.MdfyOrderTranslateStatus(orderId, status)
	return err
}

func (this *OrderController) getOrderItemById(orderItemId int64) (*models.TOrder, error) {
	// orderItem, err := models.GetOrderItemById(orderItemId)
	return nil, nil
}

func (this *OrderController) getOrderItemByProductId(productId int64, pageNo, pageSize int, where string) ([]*models.TOrderItem, int, error) {
	orderItems, totalPage, err := models.GetOrderItemByProductId(productId, pageNo, pageSize, "")
	return orderItems, totalPage, err
}

func (this *OrderController) getOrderItemByOrderId(orderId int64, pageNo, pageSize int, where string) ([]*models.TOrderItem, int, error) {
	orderItems, totalPage, err := models.GetOrderItemByOrderId(orderId, pageNo, pageSize, "")
	return orderItems, totalPage, err
}

func (this *OrderController) addOrderItem(productId int64, orderId int64) (int64, error) {
	orderItemId, err := models.AddOrderItem(productId, orderId)
	return orderItemId, err
}

func (this *OrderController) delOrderItem(orderItemId int64) error {
	err := models.DelOrderItem(orderItemId)
	return err
}
