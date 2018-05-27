package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

//订单
type TOrder struct {
	Id int64
	//订单号
	OrderNum string
	//订单类型 	OrderType == 0 普通订单 ; OrderType == 1 合作商订单 ;
	OrderType int
	//排序权重
	SortId int
	//用户Id
	UserId int64
	// //地址Id
	AddressId int64
	//备用收货地址
	Address string
	//定位
	Position string
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
	//收货人
	Receiver string
	//电话
	Phone string
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
	IsCancel  bool //是否退单
	IsRefund  bool //是否已退款。iscancel=1有效
	//客户备注
	Remark string
	//客户评价
	Comments string
	//退货原因
	CancelComments string
}

//-------------------------------基本方法------------------------------------------
//根据id获取数据实体
func GetOrderById(orderId int64) (*TOrder, error) {
	o := orm.NewOrm()
	order := new(TOrder)
	qs := o.QueryTable("t_order")
	err := qs.Filter("id", orderId).One(order)
	return order, err
}

//根据sql获取数据实体
func GetOrderBySql(excSql string) (*TOrder, error) {

	o := orm.NewOrm()
	order := new(TOrder)
	err := o.Raw(excSql).QueryRow(&order)
	return order, err
}

//新增数据实体
func AddOrder(order *TOrder) (int64, error) {
	//防止误设置id影响排序
	order.Id = 0
	order.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	//order := &TOrder{OrderTypeId: orderTypeId, UserId: userId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	orderId, err := o.Insert(order)
	return orderId, err
}

//删除数据实体
func DelOrder(orderId int64) error {

	o := orm.NewOrm()
	order := &TOrder{Id: orderId}
	_, err := o.Delete(order)
	return err
}

//批量删除数据
func DelOrders(ids string) (bool, error) {

	result := true
	sql := fmt.Sprintf("delete * from t_order where id in(%v)", ids)
	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()
	if err == nil {
		result = false
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return result, err
}

//修改数据实体
func EditOrder(order *TOrder) (int, error) {
	if order.Id == 0 {
		return 0, errors.New("id is require")
	}
	//orm模块
	ormHelper := orm.NewOrm()

	//错误对象
	num, err := ormHelper.Update(order)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//fmt.Printf("num is %v,data is %v\n", num, data)
	return int(num), err
}

//分页获取数据
func GetOrderPage(index, size int, where string) ([]*TOrder, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TOrder{}
	dataCounts := []*TOrder{}
	var sql string
	//返回数据列表
	if size == 0 {
		sql = fmt.Sprintf("select * from t_order where 1=1  %v  order by id desc", where)
	} else {
		sql = fmt.Sprintf("select * from t_order where 1=1  %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	}
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_order where 1=1 %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

//sql分页获取数据
func GetOrderPageBySql(index, size int, excSql string) ([]*TOrder, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TOrder{}
	dataCounts := []*TOrder{}
	//返回数据列表
	sql := excSql + fmt.Sprintf(" limit %v offset %v", size, size*(index-1))
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数

	count, err1 := ormHelper.Raw(excSql).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

//----------------------------扩展方法----------------------------------------

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
