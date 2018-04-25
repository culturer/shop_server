package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							注册API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *RegisterController) Post() {

	//获取数据信息
	pwd := this.Input().Get("pwd")
	tel := this.Input().Get("tel")
	name := this.Input().Get("name")
	//判断该手机号是否已经注册
	user, err := this.getUser(tel)
	if user.Id != 0 {
		beego.Info(user.Id)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 注册失败,该手机号已被注册 ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//只是打印一下信息，不需要错误处理
	if err != nil {
		beego.Info(err.Error())
	}

	//注册用户
	userId, err := this.addUser(tel, pwd, name)
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "register fail -- add user fail ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "userId": userId, "msg": "register success ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//新建User
func (this *RegisterController) addUser(tel string, password, name string) (int64, error) {
	userId, err := models.AddUser(tel, password, name)
	return userId, err
}

//获取User
func (this *RegisterController) getUser(tel string) (*models.TUser, error) {
	user, err := models.GetUserByTel(tel)
	return user, err
}
