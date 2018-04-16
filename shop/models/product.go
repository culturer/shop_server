package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//产品
type TProduct struct {
	Id int64
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
	Desc string
	//产品备注
	Msg string
}

func GetProductById(productId int64) (*TProduct, error) {
	o := orm.NewOrm()
	product := new(TProduct)
	qs := o.QueryTable("t_product")
	err := qs.Filter("id", productId).One(product)
	return product, err
}

func GetProductByType(productTypeId int64) ([]*TProduct, error) {
	products := make([]*TProduct, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_product")
	_, err := qs.Filter("product_type_id", productTypeId).All(&products)
	return products, err
}

func AddProduct(productTypeId int64, name string, count int, standardPrice float64, price float64, desc string, msg string) (int64, error) {
	o := orm.NewOrm()
	product := &TProduct{ProductTypeId: productTypeId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg}
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
