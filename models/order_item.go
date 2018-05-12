package models

import (
	"errors"
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
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
	//商品数量
	SumNum int
	//商品总价
	SumPrice float64
	//创建时间
	CreateTime string
	//商品名称
	Name string
}

//-------------------------------基本方法------------------------------------------
//根据id获取数据实体
func GetOrderItemById(OrderItemId int64) (*TOrderItem, error) {
	o := orm.NewOrm()
	OrderItem := new(TOrderItem)
	qs := o.QueryTable("t_order_item")
	err := qs.Filter("id", OrderItemId).One(OrderItem)
	return OrderItem, err
}

//根据sql获取数据实体
func GetOrderItemBySql(excSql string) (*TOrderItem, error) {

	o := orm.NewOrm()
	OrderItem := new(TOrderItem)
	err := o.Raw(excSql).QueryRow(&OrderItem)
	return OrderItem, err
}

//新增数据实体
func AddOrderItem(OrderItem *TOrderItem) (int64, error) {
	//防止误设置id影响排序
	OrderItem.Id = 0
	OrderItem.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	//OrderItem := &TOrderItem{OrderItemTypeId: OrderItemTypeId, UserId: userId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	pictureId, err := o.Insert(OrderItem)
	return pictureId, err
}

//删除数据实体
func DelOrderItem(OrderItemId int64) error {

	o := orm.NewOrm()
	OrderItem := &TOrderItem{Id: OrderItemId}
	_, err := o.Delete(OrderItem)
	return err
}

//批量删除数据
func DelOrderItems(ids string) (bool, error) {

	result := true
	sql := fmt.Sprintf("delete * from t_order_item where id in(%v)", ids)
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
func EditOrderItem(OrderItem *TOrderItem) (int, error) {
	if OrderItem.Id == 0 {
		return 0, errors.New("id is require")
	}
	//orm模块
	ormHelper := orm.NewOrm()

	//错误对象
	num, err := ormHelper.Update(OrderItem)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//fmt.Printf("num is %v,data is %v\n", num, data)
	return int(num), err
}

//分页获取数据
func GetOrderItemPage(index, size int, where string) ([]*TOrderItem, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TOrderItem{}
	dataCounts := []*TOrderItem{}
	//返回数据列表
	sql := fmt.Sprintf("select t_order_item.*,t_product.name from t_order_item left join t_product on t_order_item.product_id=t_product.id where 1=1 %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	if size == 0 {
		sql = fmt.Sprintf("select t_order_item.*,t_product.name from t_order_item left join t_product on t_order_item.product_id=t_product.id where 1=1 %v  order by id desc ", where)
	}
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_order_item where 1=1  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

//sql分页获取数据
func GetOrderItemPageBySql(index, size int, excSql string) ([]*TOrderItem, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TOrderItem{}
	dataCounts := []*TOrderItem{}
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
// func AddOrderItem(OrderItemId int64, orderId int64) (int64, error) {
// 	o := orm.NewOrm()
// 	orderItem := &TOrderItem{OrderId: orderId, OrderItemId: OrderItemId, SortId: 0}
// 	orderItemId, err := o.Insert(orderItem)
// 	return orderItemId, err
// }

// func DelOrderItem(orderItemId int64) error {
// 	o := orm.NewOrm()
// 	orderItem := &TOrderItem{Id: orderItemId}
// 	_, err := o.Delete(orderItem)
// 	return err
// }

// func GetOrderItemById(orderItemId int64) (*TOrderItem, error) {
// 	orderItem := new(TOrderItem)
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("t_order_item")
// 	err := qs.Filter("id", orderItemId).One(orderItem)
// 	return orderItem, err
// }

// func GetOrderItemByOrderItemId(OrderItemId int64, pageNo, pageSize int, where string) ([]*TOrderItem, int, error) {

// 	orderItems := make([]*TOrderItem, 0)
// 	o := orm.NewOrm()
// 	var sql string
// 	var num int64
// 	var err error
// 	if where != "" {
// 		sql = "select * from t_order_item where OrderItem_id = ? and ? order by id desc limit ? offset ?"
// 		_, err = o.Raw(sql, OrderItemId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)

// 	} else {
// 		sql = "select * from t_order_item where OrderItem_id = ? order by id desc limit ? offset ?"
// 		_, err = o.Raw(sql, OrderItemId, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
// 	}
// 	orderItems1 := make([]*TOrderItem, 0)
// 	totalNum, _ := o.Raw("select * from t_order_item where OrderItem_id = ? ", OrderItemId).QueryRows(&orderItems1)
// 	beego.Info(orderItems1)
// 	beego.Info(where)
// 	beego.Info(num)
// 	beego.Info(totalNum)
// 	mTotalNum := int(totalNum)
// 	totalPage := mTotalNum/pageSize + 1
// 	beego.Info(orderItems)
// 	return orderItems, totalPage, err

// }

// func GetOrderItemByOrderId(orderId int64, pageNo, pageSize int, where string) ([]*TOrderItem, int, error) {
// 	orderItems := make([]*TOrderItem, 0)
// 	o := orm.NewOrm()
// 	var sql string
// 	var num int64
// 	var err error
// 	if where != "" {
// 		sql = "select * from t_order_item where order_id = ? and ? order by id desc limit ? offset ?"
// 		_, err = o.Raw(sql, orderId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
// 	} else {
// 		sql = "select * from t_order_item where order_id = ? order by id desc limit ? offset ?"
// 		_, err = o.Raw(sql, orderId, pageSize, pageSize*(pageNo-1)).QueryRows(&orderItems)
// 	}
// 	orderItems1 := make([]*TOrderItem, 0)
// 	totalNum, _ := o.Raw("select * from t_order_item where order_id = ? ", orderId).QueryRows(&orderItems1)
// 	beego.Info(orderItems1)
// 	beego.Info(where)
// 	beego.Info(num)
// 	beego.Info(totalNum)
// 	mTotalNum := int(totalNum)
// 	totalPage := mTotalNum/pageSize + 1
// 	beego.Info(orderItems)
// 	return orderItems, totalPage, err
// }
