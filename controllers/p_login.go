package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"shop/models"
	"time"
)

type PLoginController struct {
	beego.Controller
}

func (c *PLoginController) Get() {
	c.TplName = "p_login.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							分销商登录				                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *PLoginController) Post() {

	//获取数据信息
	pwd := this.Input().Get("pwd")
	tel := this.Input().Get("tel")

	beego.Info(tel)
	beego.Info(pwd)

	if tel == "" || pwd == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//判断该手机号是否已经注册
	user, err := this.getUser(tel)

	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，用户不存在！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if user != nil && user.Password == pwd {
		partner, err := models.GetPartnerByUserId(user.Id)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 请先入驻成分销商！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		uid := this.GetSession("uid")
		pid := this.GetSession("pid")
		if uid == nil || pid == nil {
			//this.SetSession("uid", int(1))
			this.SetSession("uid", user.Id)
			this.SetSession("pid", partner.Id)

			beego.Info(fmt.Sprintf("uid:%v,pid:%v", this.GetSession("uid"), this.GetSession("pid"))) //this.Data["num"] = 0
		} else {
			this.SetSession("uid", user.Id)
			this.SetSession("pid", partner.Id)
			//this.Data["num"] = v.(int)
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "partner": partner, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 账号或密码错误，请重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}

//获取User
func (this *PLoginController) getUser(tel string) (*models.TUser, error) {
	user, err := models.GetUserByTel(tel)
	return user, err
}
