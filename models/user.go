package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
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
	Prov int // 0 --- 普通用户， 1 --- 管理员, 2 --- 总部用户, 3 --- 分销商
	//添加时间
	CreateTime string
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

//查询账号
func GetUserByVId(vid string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("vid", vid).One(user)
	return user, err
}

//新建用户
func AddUser(tel string, password, name string, vid string, prov int) (int64, error) {
	o := orm.NewOrm()
	user := &TUser{Password: password, Tel: tel, Prov: prov, Name: name, Vid: vid, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
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

//分页获取用户列表
func GetUserPage(pageNo, pageSize int, where string) ([]*TUser, int, error) {
	users := make([]*TUser, 0)
	o := orm.NewOrm()
	var sql string
	var sqlCount string
	//var num int64
	var err error
	if where != "" {
		sqlCount = "select * from t_user where " + where
		sql = fmt.Sprintf("select * from t_user where  %v order by id  limit %v offset %v", where, pageSize, pageSize*(pageNo-1))
		_, err = o.Raw(sql).QueryRows(&users)

	} else {
		sql = fmt.Sprintf("select * from t_user  order by id  limit %v offset %v", pageSize, pageSize*(pageNo-1))
		if pageSize == 0 {
			sql = "select * from t_user  order by id "
			sqlCount = "select * from t_user "
		}
		_, err = o.Raw(sql).QueryRows(&users)
	}
	users1 := make([]*TUser, 0)
	totalNum, _ := o.Raw(sqlCount).QueryRows(&users1)
	// beego.Info(productTypes1)
	// beego.Info(where)
	// beego.Info(num)
	// beego.Info(totalNum)
	mTotalNum := int(totalNum)
	// totalPage := mTotalNum
	// beego.Info(productTypes)
	return users, mTotalNum, err
}
