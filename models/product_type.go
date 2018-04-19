package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
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
	//排序权重
	SortId int
}

//查询分类
func GetProductTypeById(typeId int64) (*TProductType, error) {
	o := orm.NewOrm()
	productType := new(TProductType)
	qs := o.QueryTable("t_product_type")
	err := qs.Filter("id", typeId).One(productType)
	return productType, err
}

func GetProductTypeByPartnerId(partnerId int64, pageNo, pageSize int, where string) ([]*TProductType, int, error) {
	productTypes := make([]*TProductType, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_product_type where partner_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, partnerId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&productTypes)

	} else {
		sql = "select * from t_product_type where partner_id = ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, partnerId, pageSize, pageSize*(pageNo-1)).QueryRows(&productTypes)
	}
	productTypes1 := make([]*TProductType, 0)
	totalNum, _ := o.Raw("select * from t_product_type where partner_id = ? ", partnerId).QueryRows(&productTypes1)
	beego.Info(productTypes1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(productTypes)
	return productTypes, totalPage, err
}

func AddProductType(typeName string, partnerId int64) (int64, error) {
	o := orm.NewOrm()
	productType := &TProductType{TypeName: typeName, PartnerId: partnerId, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
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

func MdfyProductTypeSortId(productTypeId int64, sortId int) error {
	productType, err := GetProductTypeById(productTypeId)
	if err != nil {
		return nil
	}
	productType.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(productType)
	return err
}
