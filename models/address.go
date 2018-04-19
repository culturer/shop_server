package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//产品分类
type TAddress struct {
	Id int64
	//用户Id
	SortId int
	//排序权重
	UserId    int64
	IsDefault bool
	//国家
	Country string
	//省
	Province string
	//市
	City string
	//区
	Block string
	//街道
	Street string
	//小区
	Community string
	//具体的门牌号
	Desc string
	//手机号
	Tel string
	//N姓名
	Name string
	//创建时间
	CreateTime string
}

func AddAddress(userId int64, country string, province string, city string, block string, street string, community string, desc string, tel string, name string) (int64, error) {
	o := orm.NewOrm()
	address := &TAddress{UserId: userId, Country: country, Province: province, City: city, Block: block, Street: street, Community: community, Desc: desc, Tel: tel, Name: name, SortId: 0, CreateTime: time.Now().Format("2006-01-02 15:04:05"), IsDefault: false}
	addressId, err := o.Insert(address)
	return addressId, err
}

func DelAddress(addressId int64) error {
	o := orm.NewOrm()
	address := &TAddress{Id: addressId}
	_, err := o.Delete(address)
	return err
}

func GetAddressById(addressId int64) (*TAddress, error) {
	o := orm.NewOrm()
	address := new(TAddress)
	qs := o.QueryTable("t_address")
	err := qs.Filter("id", addressId).One(address)
	return address, err
}

func GetAddressByUserId(userId int64, pageNo, pageSize int, where string) ([]*TAddress, int, error) {
	address := make([]*TAddress, 0)
	o := orm.NewOrm()
	var sql string
	var num int64
	var err error
	if where != "" {
		sql = "select * from t_address where user_id = ? and ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, userId, where, pageSize, pageSize*(pageNo-1)).QueryRows(&address)

	} else {
		sql = "select * from t_address where user_id = ? order by id desc limit ? offset ?"
		_, err = o.Raw(sql, userId, pageSize, pageSize*(pageNo-1)).QueryRows(&address)
	}
	address1 := make([]*TOrder, 0)
	totalNum, _ := o.Raw("select * from t_address where user_id = ? ", userId).QueryRows(&address1)
	beego.Info(address1)
	beego.Info(where)
	beego.Info(num)
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	beego.Info(address)
	return address, totalPage, err

}

func MdfyAddressSort(addressId int64, sortId int) error {
	address, err := GetAddressById(addressId)
	if err != nil {
		return nil
	}
	address.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(address)
	return err
}
