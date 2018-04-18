package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (this *UserController) Post() {
	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	options, _ := strconv.Atoi(this.Input().Get("options"))
	if options == 0 {
		// [getType == 0  根据tel查询]
		// [getType == 1  根据userId查询]
		getType, _ := strconv.Atoi(this.Input().Get("getType"))
		if getType == 0 {
			tel := this.Input().Get("tel")
			user, err := this.getUserByTel(tel)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询用户异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if getType == 1 {
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			user, err := this.getUserById(userId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询用户异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	if options == 1 {
		tel := this.Input().Get("tel")
		password := this.Input().Get("password")
		userId, err := this.addUser(tel, password)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 新增用户异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "userId": userId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
		err := this.delUser(userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除用户异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 删除用户成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		// [mdfyType == 0  修改手机号]
		// [mdfyType == 1  修改密码]
		// [mdfyType == 2  修改第三方Id]
		// [mdfyType == 3  修改权限]
		mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
		if mdfyType == 0 {
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			tel := this.Input().Get("tel")
			err := this.mdfyTel(userId, tel)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改手机号异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改手机号成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 1 {
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			password := this.Input().Get("password")
			err := this.mdfyPassword(userId, password)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改密码异常,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改密码成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 2 {
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			vid := this.Input().Get("vid")
			err := this.mdfyVid(userId, vid)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改第三方Id,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改第三方Id成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 3 {
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			prov, _ := strconv.Atoi(this.Input().Get("prov"))
			err := this.mdfyProv(userId, prov)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改第权限,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改权限成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}
}

//查询账号
func (this *UserController) getUserById(userId int64) (*models.TUser, error) {
	user, err := models.GetUserById(userId)
	return user, err
}

//查询账号
func (this *UserController) getUserByTel(tel string) (*models.TUser, error) {
	user, err := models.GetUserByTel(tel)
	return user, err
}

//新建用户
func (this *UserController) addUser(tel string, password string) (int64, error) {
	userId, err := models.AddUser(tel, password)
	return userId, err
}

//删除账号
func (this *UserController) delUser(userId int64) error {
	err := models.DelUser(userId)
	return err
}

//修改手机号
func (this *UserController) mdfyTel(userId int64, tel string) error {
	err := models.MdfyTel(userId, tel)
	return err
}

//修改密码
func (this *UserController) mdfyPassword(userId int64, password string) error {
	err := models.MdfyPassword(userId, password)
	return err
}

//修改第三方Id
func (this *UserController) mdfyVid(userId int64, vid string) error {
	err := models.MdfyVid(userId, vid)
	return err
}

//修改权限
func (this *UserController) mdfyProv(userId int64, prov int) error {
	err := models.MdfyProv(userId, prov)
	return err
}
