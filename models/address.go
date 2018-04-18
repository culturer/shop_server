package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"time"
)

//产品分类
type TAddress struct {
	Id int64
	//用户Id
	Sort_Id int
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
	address := &TAddress{UserId: userId, Country: country, Province: province, City: city, Block: block, Street: street, Community: community, Desc: desc, Tel: tel, Name: name, Sort_Id: 0, CreateTime: time.Now().Format("2006-01-02 15:04:05"), IsDefault: false}
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

func GetAddressByUserId(userId int64) ([]*TAddress, error) {
	address := make([]*TAddress, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_address")
	_, err := qs.Filter("user_id", userId).All(&address)
	return address, err
}

func MdfyAddressSort(addressId int64, sort_Id int) error {
	address, err := GetAddressById(addressId)
	if err != nil {
		return nil
	}
	address.Sort_Id = sort_Id
	o := orm.NewOrm()
	_, err = o.Update(address)
	return err
}
