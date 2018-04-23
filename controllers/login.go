package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {

	//获取数据信息
	pwd := this.Input().Get("pwd")
	tel := this.Input().Get("tel")

	//判断该手机号是否已经注册
	user, err := this.getUser(tel)

	if err != nil {
		beego.Info(err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，账号或密码错误，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if user != nil && user.Password == pwd {
		this.SetSession("uid", int(1))
		this.Data["json"] = map[string]interface{}{"status": 200, "userId": user.Id, "msg": "login success ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，账号或密码错误，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//获取User
func (this *LoginController) getUser(tel string) (*models.TUser, error) {
	user, err := models.GetUserByTel(tel)
	return user, err
}
