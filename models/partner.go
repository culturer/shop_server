package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//分销商
type TPartner struct {
	Id int64
	//用户Id
	UserId int64
	//分析商名称
	PartnerName string
	//地址
	AddressId int64
	//积分
	Credits int
	//权限
	Pro int
}

func AddPartner(userId int64, partnerName string, addressId int64) (int64, error) {
	o := orm.NewOrm()
	partner := &TPartner{UserId: userId, PartnerName: partnerName, AddressId: addressId, Credits: 0, Pro: 0}
	partnerId, err := o.Insert(partner)
	return partnerId, err
}

func DelPartner(partnerId int64) error {
	o := orm.NewOrm()
	partner := &TPartner{Id: partnerId}
	_, err := o.Delete(partner)
	return err
}

func GetPartnerById(partnerId int64) (*TPartner, error) {
	partner := new(TPartner)
	o := orm.NewOrm()
	qs := o.QueryTable("t_partner")
	err := qs.Filter("id", partnerId).One(partner)
	return partner, err
}

func GetPartners() ([]*TPartner, error) {
	partners := make([]*TPartner, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_partner")
	_, err := qs.All(&partners)
	return partners, err
}

func MdfyPartnerName(partnerId int64, partnerName string) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.PartnerName = partnerName
	o := orm.NewOrm()
	_, err = o.Update(partner)
	return err
}

func MdfyPartnerAddress(partnerId int64, addressId int64) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.AddressId = addressId
	o := orm.NewOrm()
	_, err = o.Update(partner)
	return err
}

func MdfyPartnerCredits(partnerId int64, credits int) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.Credits = credits
	o := orm.NewOrm()
	_, err = o.Update(partner)
	return err
}

func MdfyPartnerPro(partnerId int64, pro int) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.Pro = pro
	o := orm.NewOrm()
	_, err = o.Update(partner)
	return err
}
