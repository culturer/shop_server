package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"strconv"
	"strings"
	"time"
)

//订单
type TOrder struct {
	Id int64
	//用户Id
	SortId int
	//排序权重
	UserId int64
	//地址Id
	AddressId int64
	//备用收货地址
	Address string
	//物流状态
	TranslateStatus string
	//付款方式
	PayType string
	//实际金额
	RealPrice float64
	//应付金额
	ShouldPrice float64
	//金额注释
	PriceMsg string
	//分销商
	PartnerId int64
	//下单时间
	CreateTime string
	//订单状态 0 --- 待付款， 1 --- 已付款 ， 2 --- 待发货 ， 3 --- 已发货 ， 4 --- 待签收 ， 5 --- 已签收 ， 6 --- 已评价
	OrderStatus int
	//订单状态，新
	IsPay     bool //付款
	IsDlivery bool //发货
	IsSign    bool //签收
	IsCash    bool //货到付款
	IsComment bool //评论
}

func AddOrder(userId int64, addressId int64, address string, payType string, realPrice float64, shouldPrice float64, priceMsg string, partnerId int64) (int64, error) {
	o := orm.NewOrm()
	order := &TOrder{UserId: userId, AddressId: addressId, Address: address, PayType: payType, PartnerId: partnerId, CreateTime: time.Now().Format("2006-01-02 15:04:05"), OrderStatus: 0, SortId: 0}
	userId, err := o.Insert(order)
	return userId, err
}

func DelOrder(orderId int64) error {
	o := orm.NewOrm()
	order := &TOrder{Id: orderId}
	_, err := o.Delete(order)
	return err
}

//查询

func GetOrderById(orderId int64) (*TOrder, error) {
	order := new(TOrder)
	o := orm.NewOrm()
	qs := o.QueryTable("t_order")
	err := qs.Filter("id", orderId).One(order)
	return order, err
}

func GetOrderByUserId(userId int64, pageNo, pageSize int, where string) ([]*TOrder, int, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_order where user_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, userId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orders)

	} else {
		sql = "select * from t_order where user_id = ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, userId, pageSize, pageSize*(pageNo-1)).QueryRows(&orders)
	}
	orders1 := make([]*TOrder, 0)
	totalNum, _ := o.Raw("select * from t_order where user_id = ? ", userId).QueryRows(&orders1)
	beego.Info(orders1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(orders)
	return orders, totalPage, err
}

func GetOrderByUserIdS(userId int64, orderStatus int, pageNo, pageSize int) ([]*TOrder, int, error) {

	where := strings.Join([]string{" order_status = ", strconv.Itoa(orderStatus)}, "")
	orders, totalPage, err := GetOrderByUserId(userId, pageNo, pageSize, where)
	return orders, totalPage, err

}

func GetOrderByPartnerId(partnerId int64, pageNo, pageSize int, where string) ([]*TOrder, int, error) {
	orders := make([]*TOrder, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_order where partner_id = ? and ? order by id desc limit ? offset ?"
		num, err = o.Raw(sql, partnerId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orders)

	} else {
		sql = "select * from t_order where partner_id = ? order by id desc limit ? offset ?"
		num, err = o.Raw(sql, partnerId, pageSize, pageSize*(pageNo-1)).QueryRows(&orders)
	}
	orders1 := make([]*TOrder, 0)
	totalNum, _ := o.Raw("select * from t_order where partner_id = ? ", partnerId).QueryRows(&orders1)
	beego.Info(orders1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(orders)
	return orders, totalPage, err
}

func GetOrderByIdPartnerIdS(partnerId int64, pageNo, pageSize, orderStatus int) ([]*TOrder, int, error) {

	where := strings.Join([]string{"order_status = ", strconv.Itoa(orderStatus)}, "")
	orders, totalPage, err := GetOrderByPartnerId(partnerId, pageNo, pageSize, where)
	return orders, totalPage, err

}

//修改

func MdfyOrderStatus(orderId int64, orderStatus int, payType string) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.OrderStatus = orderStatus
	order.PayType = payType
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func OrderIsPay(orderId int64, isPay bool) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.IsPay = isPay
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func OrderIsDlivery(orderId int64, isDlivery bool) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.IsDlivery = isDlivery
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func OrderIsSign(orderId int64, isSign bool) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.IsSign = isSign
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func OrderIsCash(orderId int64, isCash bool) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.IsCash = isCash
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func OrderIsComment(orderId int64, isComment bool) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.IsComment = isComment
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

func MdfyOrderSort(orderId int64, sortId int) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}

func MdfyOrderTranslateStatus(orderId int64, status string) error {
	order, err := GetOrderById(orderId)
	if err != nil {
		return nil
	}
	order.TranslateStatus = status
	o := orm.NewOrm()
	_, err = o.Update(order)
	return err
}
