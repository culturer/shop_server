package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//产品
type TProduct struct {
	Id int64
	//排序权重
	SortId int
	UserId int64
	//产品分类
	ProductTypeId int64
	//产品名称
	Name string
	//库存量
	Count int
	//价格
	Price float64
	//成本价
	StandardPrice float64
	//产品描述
	Desc string `orm:"size(5000)"`
	//产品备注
	Msg string `orm:"size(2000)"`
	//创建时间
	CreateTime string
}

func GetProductById(productId int64) (*TProduct, error) {
	o := orm.NewOrm()
	product := new(TProduct)
	qs := o.QueryTable("t_product")
	err := qs.Filter("id", productId).One(product)
	return product, err
}

func GetProductByType(productTypeId int64, pageNo, pageSize int, where string) ([]*TProduct, int, error) {
	products := make([]*TProduct, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_product where product_type_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, productTypeId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&products)

	} else {
		sql = fmt.Sprintf("select * from t_product where product_type_id = %v order by id desc limit %v offset %v", productTypeId, pageSize, pageSize*(pageNo-1))
		_, err = o.Raw(sql).QueryRows(&products)
	}
	products1 := make([]*TProduct, 0)
	totalNum, _ := o.Raw("select * from t_product where product_type_id = ? ", productTypeId).QueryRows(&products1)
	beego.Info(sql)
	beego.Info(products1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(products)
	return products, totalPage, err
}

func AddProduct(productTypeId int64, userId int64, name string, count int, standardPrice float64, price float64, desc string, msg string) (int64, error) {
	o := orm.NewOrm()
	product := &TProduct{ProductTypeId: productTypeId, UserId: userId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	pictureId, err := o.Insert(product)
	return pictureId, err
}

func DelProduct(productId int64) error {
	o := orm.NewOrm()
	product := &TProduct{Id: productId}
	_, err := o.Delete(product)
	return err
}

func MdfyType(productId int64, productTypeId int64) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.ProductTypeId = productTypeId
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyName(productId int64, name string) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.Name = name
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyCount(productId int64, count int) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.Count = count
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyPrice(productId int64, price float64) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.Price = price
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyStandardPrice(productId int64, standardPrice float64) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.StandardPrice = standardPrice
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyDesc(productId int64, desc string) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.Desc = desc
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyMsg(productId int64, msg string) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.Msg = msg
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}

func MdfyProductSort(productId int64, sortId int) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(product)
	return err
}
