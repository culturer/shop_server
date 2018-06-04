package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"errors"
	"fmt"
	"time"
)

//广告
type TAdvertise struct {
	Id int64
	//标题
	Title string
	//内容
	Content string `orm:"type(text);null"`
	//浏览次数
	Count int
	//关联的产品Id，方便直接打开产品信息
	ProductId int64
	//创建时间
	CreateTime string
	//置顶
	IsTop bool
}

//新增数据实体
func AddAdvertise(advertise *TAdvertise) (int64, error) {
	//防止误设置id影响排序
	advertise.Id = 0
	advertise.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	advertiseId, err := o.Insert(advertise)
	return advertiseId, err
}

func DelAdvertise(advertiseId int64) error {
	o := orm.NewOrm()
	advertise := &TAdvertise{Id: advertiseId}
	_, err := o.Delete(advertise)
	return err
}

func GetAdvertiseById(advertiseId int64) (*TAdvertise, error) {
	o := orm.NewOrm()
	advertise := new(TAdvertise)
	qs := o.QueryTable("t_advertise")
	err := qs.Filter("id", advertiseId).One(advertise)
	return advertise, err
}

func GetAdvertiseByPage(index, size int, where string) ([]*TAdvertise, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TAdvertise{}
	dataCounts := []*TAdvertise{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_advertise %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_advertise  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}

func MdfyAdvertise(advertise *TAdvertise) (int, error) {
	if advertise.Id == 0 {
		return 0, errors.New("id is require")
	}
	//orm模块
	ormHelper := orm.NewOrm()

	//错误对象
	num, err := ormHelper.Update(advertise)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//fmt.Printf("num is %v,data is %v\n", num, data)
	return int(num), err
}
