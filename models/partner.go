package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
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
	//备注
	Desc string
	//添加时间
	Add_time string
}

func AddPartner(userId int64, partnerName string, address, position, desc string) (int64, error) {
	o := orm.NewOrm()
	partner := &TPartner{UserId: userId, PartnerName: partnerName, Address: address, Position: position, Desc: desc, Credits: 0, Pro: 0, SortId: 0, Add_time: time.Now().Format("2006-01-02 15:04:05")}
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

func GetPartnerByUserId(userId int64) (*TPartner, error) {
	partner := new(TPartner)
	o := orm.NewOrm()
	qs := o.QueryTable("t_partner")
	err := qs.Filter("user_id", userId).One(partner)
	return partner, err
}

func GetPartners(pageNo, pageSize int) ([]*TPartner, int, error) {
	partners := make([]*TPartner, 0)
	o := orm.NewOrm()

	sql := "select * from t_partner  order by id desc limit ? offset ?"
	if pageSize == 0 {
		sql = "select * from t_partner  order by id desc "
	}
	totalNum, err := o.Raw(sql, pageSize, pageSize*(pageNo-1)).QueryRows(&partners)
	totalNum, _ = o.Raw("select * from t_partner ").QueryRows(new([]TPartner))
	beego.Info(totalNum)
	mTotalNum := int(totalNum)
	//totalPage := mTotalNum/pageSize + 1
	return partners, mTotalNum, err

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
