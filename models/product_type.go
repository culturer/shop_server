package models

import (
	"errors"
	"fmt"
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
	//PartnerId int64
	//创建时间
	CreateTime string
	//排序权重
	SortId int64
}

//-------------------------------基本方法------------------------------------------
//根据id获取数据实体
func GetProductTypeById(productId int64) (*TProductType, error) {
	o := orm.NewOrm()
	productType := new(TProductType)
	qs := o.QueryTable("t_product_type")
	err := qs.Filter("id", productId).One(productType)
	return productType, err
}

//根据sql获取数据实体
func GetProductTypeBySql(excSql string) (*TProductType, error) {

	o := orm.NewOrm()
	productType := new(TProductType)
	err := o.Raw(excSql).QueryRow(&productType)
	return productType, err
}

//新增数据实体
func AddProductType(productType *TProductType) (int64, error) {
	//防止误设置id影响排序
	productType.Id = 0
	productType.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	//product := &TProduct{ProductTypeId: productTypeId, UserId: userId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	productTypeId, err := o.Insert(productType)
	return productTypeId, err
}

//删除数据实体
func DelProductType(productTypeId int64) error {

	o := orm.NewOrm()
	productType := &TProductType{Id: productTypeId}
	_, err := o.Delete(productType)
	return err
}

//批量删除数据
func DelProductTypes(ids string) (bool, error) {

	result := true
	sql := fmt.Sprintf("delete * from t_product_type id in(%v)", ids)
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
func EditProductType(productType *TProductType) (int, error) {
	if productType.Id == 0 {
		return 0, errors.New("id is require")
	}
	//orm模块
	ormHelper := orm.NewOrm()

	//错误对象
	num, err := ormHelper.Update(productType)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//fmt.Printf("num is %v,data is %v\n", num, data)
	return int(num), err
}

//分页获取数据
func GetProductTypePage(index, size int, where string) ([]*TProductType, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TProductType{}
	dataCounts := []*TProductType{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_product_type %v  order by sort_id desc limit %v offset %v", where, size, size*(index-1))
	if size == 0 {
		sql = "select * from t_product_type order by sort_id desc"
	}
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		beego.Info(sql)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_product_type  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		beego.Info(err1)
	}
	return data, int(count), err
}

//sql分页获取数据
func GetProductTypePageBySql(index, size int, excSql string) ([]*TProductType, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TProductType{}
	dataCounts := []*TProductType{}
	//返回数据列表
	sql := excSql + fmt.Sprintf(" limit %v offset %v", size, size*(index-1))
	if size == 0 {
		sql = excSql
	}
	_, err := ormHelper.Raw(sql).QueryRows(data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数

	count, err1 := ormHelper.Raw(excSql).QueryRows(dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

//----------------------------扩展方法----------------------------------------
//查询分类
// func GetProductTypeById(typeId int64) (*TProductType, error) {
// 	o := orm.NewOrm()
// 	productType := new(TProductType)
// 	qs := o.QueryTable("t_product_type")
// 	err := qs.Filter("id", typeId).One(productType)
// 	return productType, err
// }

// func GetProductTypePage(pageNo, pageSize int, where string) ([]*TProductType, int, error) {
// 	productTypes := make([]*TProductType, 0)
// 	o := orm.NewOrm()
// 	var sql string
// 	var num int64
// 	var err error
// 	if where != "" {

// 		sql = fmt.Sprintf("select * from t_product_type where  %v order by sort_id  limit %v offset %v", where, pageSize, pageSize*(pageNo-1))
// 		_, err = o.Raw(sql).QueryRows(&productTypes)

// 	} else {
// 		sql = fmt.Sprintf("select * from t_product_type  order by sort_id  limit %v offset %v", pageSize, pageSize*(pageNo-1))
// 		if pageSize == 0 {
// 			sql = "select * from t_product_type  order by sort_id "
// 		}
// 		_, err = o.Raw(sql).QueryRows(&productTypes)
// 	}
// 	productTypes1 := make([]*TProductType, 0)
// 	totalNum, _ := o.Raw("select * from t_product_type ").QueryRows(&productTypes1)
// 	beego.Info(productTypes1)
// 	beego.Info(where)
// 	beego.Info(num)
// 	beego.Info(totalNum)
// 	mTotalNum := int(totalNum)
// 	totalPage := mTotalNum
// 	beego.Info(productTypes)
// 	return productTypes, totalPage, err
// }

// func AddProductType(typeName string, sortId int64) (int64, error) {
// 	o := orm.NewOrm()
// 	productType := &TProductType{TypeName: typeName, SortId: sortId, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
// 	productTypeId, err := o.Insert(productType)
// 	return productTypeId, err
// }

// func DelProductType(productTypeId int64) error {
// 	o := orm.NewOrm()
// 	productType := &TProductType{Id: productTypeId}
// 	_, err := o.Delete(productType)
// 	return err
// }

func MdfyPartner(productTypeId int64, sortId int64) error {
	productType, err := GetProductTypeById(productTypeId)
	if err != nil {
		return nil
	}
	productType.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(productType)
	return err
}

func MdfyProductTypeSortId(productTypeId int64, sortId int64) error {
	productType, err := GetProductTypeById(productTypeId)
	if err != nil {
		return nil
	}
	productType.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(productType)
	return err
}
