package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分销商
type TPartner struct {
	Id int64
	//用户Id
	UserId int64
	//分析商名称
	PartnerName string
	//地址
	Address string
	//积分
	Credits int
	//权限
	Pro int
	//排序权重
	SortId int
	//位置
	Position string
	//添加时间
	Add_time string
}

func AddPartner(userId int64, partnerName string, address string) (int64, error) {
	o := orm.NewOrm()
	partner := &TPartner{UserId: userId, PartnerName: partnerName, Address: address, Credits: 0, Pro: 0, SortId: 0}
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

func GetPartners(pageNo, pageSize int) ([]*TPartner, int, error) {
	partners := make([]*TPartner, 0)
	o := orm.NewOrm()

	sql := "select * from t_partner  order by id desc limit ? offset ?"
	totalNum, err := o.Raw(sql, pageSize, pageSize*(pageNo-1)).QueryRows(&partners)

	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	totalPage := mTotalNum/pageSize + 1
	return partners, totalPage, err

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

func MdfyPartnerAddress(partnerId int64, address string) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.Address = address
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

func MdfyPartnerSort(partnerId int64, sortId int) error {
	partner, err := GetPartnerById(partnerId)
	if err != nil {
		return nil
	}
	partner.SortId = sortId
	o := orm.NewOrm()
	_, err = o.Update(partner)
	return err
}
