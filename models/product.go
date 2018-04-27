package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"errors"
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
	//产品分类
	PartnerId int64
	//产品名称
	Name string
	//库存量
	Count int
	//价格
	Price float64
	//成本价
	StandardPrice float64
	//产品描述
	Desc string `orm:"type(text);null"`
	//产品备注
	Msg string `orm:"type(text);null"`

	//创建时间
	CreateTime string
}

//-------------------------------基本方法------------------------------------------
//根据id获取数据实体
func GetProductById(productId int64) (*TProduct, error) {
	o := orm.NewOrm()
	product := new(TProduct)
	qs := o.QueryTable("t_product")
	err := qs.Filter("id", productId).One(product)
	return product, err
}

//根据sql获取数据实体
func GetProductBySql(excSql string) (*TProduct, error) {

	o := orm.NewOrm()
	product := new(TProduct)
	err := o.Raw(excSql).QueryRow(&product)
	return product, err
}

//新增数据实体
func AddProduct(product *TProduct) (int64, error) {
	//防止误设置id影响排序
	product.Id = 0
	product.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	//product := &TProduct{ProductTypeId: productTypeId, UserId: userId, Name: name, Count: count, StandardPrice: standardPrice, Price: price, Desc: desc, Msg: msg, CreateTime: time.Now().Format("2006-01-02 15:04:05"), SortId: 0}
	pictureId, err := o.Insert(product)
	return pictureId, err
}

//删除数据实体
func DelProduct(productId int64) error {

	o := orm.NewOrm()
	product := &TProduct{Id: productId}
	_, err := o.Delete(product)
	return err
}

//批量删除数据
func DelProducts(ids string) (bool, error) {

	result := true
	sql := fmt.Sprintf("delete * from t_product id in(%v)", ids)
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
func EditProduct(product *TProduct) (int, error) {
	if product.Id == 0 {
		return 0, errors.New("id is require")
	}
	//orm模块
	ormHelper := orm.NewOrm()

	//错误对象
	num, err := ormHelper.Update(product)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//fmt.Printf("num is %v,data is %v\n", num, data)
	return int(num), err
}

//分页获取数据
func GetProductPage(index, size int, where string) ([]*TProduct, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TProduct{}
	dataCounts := []*TProduct{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_product %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_product  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

//sql分页获取数据
func GetProductPageBySql(index, size int, excSql string) ([]*TProduct, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TProduct{}
	dataCounts := []*TProduct{}
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
//根据商品类别获取分页数据
func GetProductByType(productTypeId int64, pageNo, pageSize int, where string) ([]*TProduct, int, error) {
	products := make([]*TProduct, 0)
	o := orm.NewOrm()
	var sql string
	//var num int64
	var err error
	if where != "" {
		sql = "select * from t_product where product_type_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, productTypeId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&products)

	} else {
		sql = fmt.Sprintf("select * from t_product where product_type_id = %v order by id desc limit %v offset %v", productTypeId, pageSize, pageSize*(pageNo-1))
		if productTypeId == 0 {
			sql = fmt.Sprintf("select * from t_product order by id desc limit %v offset %v", pageSize, pageSize*(pageNo-1))

		}
		_, err = o.Raw(sql).QueryRows(&products)
	}
	products1 := make([]*TProduct, 0)
	sqlCount := fmt.Sprintf("select * from t_product where product_type_id = %v ", productTypeId)
	if productTypeId == 0 {
		sqlCount = "select * from t_product"

	}
	totalNum, _ := o.Raw(sqlCount).QueryRows(&products1)
	//beego.Info(sql)
	// beego.Info(products1)
	// beego.Info(where)
	// beego.Info(num)
	// beego.Info(totalNum)
	// mTotalNum := int(totalNum)
	// totalPage := mTotalNum/pageSize + 1
	// beego.Info(products)
	return products, int(totalNum), err
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
