package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//用户
type TUser struct {
	Id int64
	//姓名
	Name string
	//手机号
	Tel string
	//库存数量
	Password string
	//第三方账号id
	Vid string
	//权限
	Prov int
}

//查询账号
func GetUserById(userId int64) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("id", userId).One(user)
	return user, err
}

//查询账号
func GetUserByTel(tel string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("tel", tel).One(user)
	return user, err
}

//新建用户
func AddUser(tel string, password string) (int64, error) {
	o := orm.NewOrm()
	user := &TUser{Password: password, Tel: tel}
	userId, err := o.Insert(user)
	return userId, err
}

//删除账号
func DelUser(userId int64) error {
	o := orm.NewOrm()
	user := &TUser{Id: userId}
	_, err := o.Delete(user)
	return err
}

//修改手机号
func MdfyTel(userId int64, tel string) error {
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	user.Tel = tel
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

//修改密码
func MdfyPassword(userId int64, password string) error {
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	user.Password = password
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

//修改第三方Id
func MdfyVid(userId int64, vid string) error {
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	user.Vid = vid
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

//修改权限
func MdfyProv(userId int64, prov int) error {
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	user.Prov = prov
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}
