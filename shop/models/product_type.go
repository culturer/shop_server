package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//产品分类
type TProductType struct {
	Id int64
	//类别名称
	TypeName string
	//分销商Id
	PartnerId int64
	//创建时间
	CreateTime string
}

//查询分类
func GetProductTypeById(typeId int64) (*TProductType, error) {
	o := orm.NewOrm()
	productType := new(TProductType)
	qs := o.QueryTable("t_product_type")
	err := qs.Filter("id", typeId).One(productType)
	return productType, err
}

func GetProductTypeByPartnerId(partnerId int64) ([]*TProductType, error) {
	productTypes := make([]*TProductType, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_product_type")
	_, err := qs.Filter("partner_id", partnerId).All(productTypes)
	return productTypes, err
}

func AddProductType(typeName string, partnerId int64) (int64, error) {
	o := orm.NewOrm()
	productType := &TProductType{TypeName: typeName, PartnerId: partnerId, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	productTypeId, err := o.Insert(productType)
	return productTypeId, err
}

func DelProductType(productTypeId int64) error {
	o := orm.NewOrm()
	productType := &TProductType{Id: productTypeId}
	_, err := o.Delete(productType)
	return err
}

func MdfyPartner(productTypeId int64, partnerId int64) error {
	productType, err := GetProductTypeById(productTypeId)
	if err != nil {
		return nil
	}
	productType.PartnerId = partnerId
	o := orm.NewOrm()
	_, err = o.Update(productType)
	return err
}
