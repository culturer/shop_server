package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	var page string
	this.Ctx.Input.Bind(&page, "page")
	if page == "index" {
		this.TplName = "index.html"
	} else {
		this.TplName = "login.html"
	}

}

func (this *LoginController) Post() {

	// [options == 0  客户登录]
	// [options == 1  管理员登录]
	// [options == 2  微信登录]
	options, _ := strconv.Atoi(this.Input().Get("options"))

	if options == 0 || options == 1 {

		//获取数据信息
		pwd := this.Input().Get("pwd")
		tel := this.Input().Get("tel")

		//判断该手机号是否已经注册
		user, err := this.getUser(tel)

		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，用户不存在，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if user != nil && user.Password == pwd {
			//this.SetSession("uid", int(64))
			uid := this.GetSession("uid")
			if uid == nil {
				//this.SetSession("uid", int(1))
				this.SetSession("uid", user.Id)

				beego.Info(fmt.Sprintf("uid:%v", this.GetSession("uid"))) //this.Data["num"] = 0
			} else {
				this.SetSession("uid", user.Id)
				//this.Data["num"] = v.(int)
			}

			//返回Json数据，提供给Android 或者 IOS APP使用
			if options == 0 {
				this.Data["json"] = map[string]interface{}{"status": 200, "userId": user.Id, "msg": "login success ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			//判断是否是管理员登录
			if options == 1 {
				if user.Prov == 1 {
					this.Data["json"] = map[string]interface{}{"status": 200, "userId": user.Id, "msg": "login success ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，账号或密码错误，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "userId": user.Id, "msg": "login success ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	} else if options == 2 {
		vid := this.Input().Get("vId")
		user, err := this.getUserByVid(vid)
		if err != nil || user == nil || user.Tel == "" || user.Password == "" {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 登录失败，账号或密码错误，请检查后重写登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		//存session
		uid := this.GetSession("uid")
		if uid == nil {
			//this.SetSession("uid", int(1))
			this.SetSession("uid", user.Id)
			beego.Info(fmt.Sprintf("uid:%v", this.GetSession("uid"))) //this.Data["num"] = 0
		} else {
			this.SetSession("uid", user.Id)
			//this.Data["num"] = v.(int)
		}
		//返回userId
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

func (this *LoginController) getUserByVid(vid string) (*models.TUser, error) {
	user, err := models.GetUserByVId(vid)
	return user, err
}
