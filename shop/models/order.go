package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"time"
)

//订单
type TOrder struct {
	Id int64
	//用户Id
	UsertId int64
	//地址Id
	AddressId int64
	//分销商
	PartnerId int64
	//下单时间
	CreateTime string
	//订单状态 0 --- 待付款， 1 --- 已付款 ， 2 --- 待发货 ， 3 --- 已发货 ， 4 --- 待签收 ， 5 --- 已签收 ， 6 --- 已评价
	OrderStatus int
}

func AddOrder(userId int64, addressId int64, partnerId int64) (int64, error) {
	o := orm.NewOrm()
	order := &TOrder{UsertId: userId, AddressId: addressId, PartnerId: partnerId, CreateTime: time.Now().Format("2006-01-02 15:04:05"), OrderStatus: 0}
	userId, err := o.Insert(order)
	return userId, err
}

func DelOrder(orderId int64) error {
	o := orm.NewOrm()
	order := &TOrder{Id: orderId}
	_, err := o.Delete(order)
	return err
}

func GetOrderById(orderId int64) (*TOrder, error) {
	order := new(TOrder)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	err := qs.Filter("id", orderId).One(order)
	return order, err
}

func GetOrderByUserId(userId int64) ([]*TOrder, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	_, err := qs.Filter("user_id", userId).All(&orders)
	return orders, err
}

func GetOrderByUserIdS(userId int64, orderStatus int) ([]*TOrder, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	_, err := qs.Filter("user_id", userId).Filter("order_status", orderStatus).All(&orders)
	return orders, err
}

func GetOrderByPartnerId(partnerId int64) ([]*TOrder, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	_, err := qs.Filter("partner_id", partnerId).All(&orders)
	return orders, err
}

func GetOrderByIdPartnerIdS(partnerId int64, orderStatus int) ([]*TOrder, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	_, err := qs.Filter("partner_id", partnerId).Filter("order_status", orderStatus).All(&orders)
	return orders, err
}

func MdfyOrderStatus(orderId int64, orderStatus int) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.OrderStatus = orderStatus
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func MdfyOrderAddress(orderId int64, addressId int64) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.AddressId = addressId
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}
