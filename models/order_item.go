package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//订单商品项
type TOrderItem struct {
	Id int64
	//订单Id
	OrderId int64
	//商品Id
	ProductId int64
	//排序权重
	SortId int
}

func AddOrderItem(productId int64, orderId int64) (int64, error) {
	o := orm.NewOrm()
	orderItem := &TOrderItem{OrderId: orderId, ProductId: productId, SortId: 0}
	orderItemId, err := o.Insert(orderItem)
	return orderItemId, err
}

func DelOrderItem(orderItemId int64) error {
	o := orm.NewOrm()
	orderItem := &TOrderItem{Id: orderItemId}
	_, err := o.Delete(orderItem)
	return err
}

func GetOrderItemById(orderItemId int64) (*TOrder, error) {
	orderItem := new(TOrder)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order_item")
	err := qs.Filter("id", orderItemId).One(orderItem)
	return orderItem, err
}

func GetOrderItemByProductId(productId int64) ([]*TOrderItem, error) {
	orderItems := make([]*TOrderItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order_item")
	_, err := qs.Filter("productr_id", productId).All(&orderItems)
	return orderItems, err
}

func GetOrderItemByOrderId(orderId int64) ([]*TOrderItem, error) {
	orderItems := make([]*TOrderItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order_item")
	_, err := qs.Filter("order_id", orderId).All(&orderItems)
	return orderItems, err
}
