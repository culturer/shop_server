package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"time"
)

//产品
type TProduct struct {
	Id int64
	//排序权重
	SortId int
	//产品分类
	ParentId int64
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

func GetProductByType(parentId int64) ([]*TProduct, error) {
	products := make([]*TProduct, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_product")
	_, err := qs.Filter("product_type_id", parentId).All(&products)
	return products, err
}

func AddProduct(parentId int64, name string, count int, standardPrice float64, price float64, desc string, msg string) (int64, error) {
	o := orm.NewOrm()
	product := &TProduct{ParentId: parentId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	pictureId, err := o.Insert(product)
	return pictureId, err
}

func DelProduct(productId int64) error {
	o := orm.NewOrm()
	product := &TProduct{Id: productId}
	_, err := o.Delete(product)
	return err
}

func MdfyType(productId int64, parentId int64) error {
	product, err := GetProductById(productId)
	if err != nil {
		return nil
	}
	product.ParentId = parentId
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
