package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func GetOrderItemByProductId(productId int64, pageNo, pageSize int, where string) ([]*TOrderItem, int, error) {

	orderItems := make([]*TOrderItem, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_order_item where product_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, productId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)

	} else {
		sql = "select * from t_order_item where product_id = ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, productId, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
	}
	orderItems1 := make([]*TOrderItem, 0)
	totalNum, _ := o.Raw("select * from t_order_item where product_id = ? ", productId).QueryRows(&orderItems1)
	beego.Info(orderItems1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(orderItems)
	return orderItems, totalPage, err

}

func GetOrderItemByOrderId(orderId int64, pageNo, pageSize int, where string) ([]*TOrderItem, int, error) {
	orderItems := make([]*TOrderItem, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_order_item where order_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, orderId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
	} else {
		sql = "select * from t_order_item where order_id = ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, orderId, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
	}
	orderItems1 := make([]*TOrderItem, 0)
	totalNum, _ := o.Raw("select * from t_order_item where order_id = ? ", orderId).QueryRows(&orderItems1)
	beego.Info(orderItems1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(orderItems)
	return orderItems, totalPage, err
}
