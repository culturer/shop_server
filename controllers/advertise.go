package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"
	"time"
)

type AdvertiseController struct {
	beego.Controller
}

func (this *AdvertiseController) Get() {
	var page string
	this.Ctx.Input.Bind(&page, "page")
	if page == "cover_list" {
		this.TplName = "cover_list.html"
	} else if page == "advertise_list" {
		this.TplName = "advertise_list.html"
	} else if page == "product_type_list" {
		this.TplName = "product_type_list.html"
	} else if page == "product_edit" {
		this.TplName = "product_edit.html"
	}
}

//////////////////////////////////////////////////////////////////////
//																	//
//							广告API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *AdvertiseController) Post() {

	//获取数据信息
	options := this.Input().Get("options")

	if options != "" {
		switch options {
		case "isCorver":
			//修改首页轮播
			this.isCorver()
		case "getCovers":
			//获取广告轮播
			this.getCovers()
		case "addadvertise":
			//增加广告
			this.addAdvertise()
		case "deladvertise":
			//删除广告
			this.delAdvertise()
		case "mdfyAdvertise":
			//修改广告
			this.mdfyAdvertise()
		case "getAdvertiseByPage":
			this.getAdvertiseByPage()
		}
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "register success ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

func (this *AdvertiseController) isCorver() {

	productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
	pictureId, _ := strconv.ParseInt(this.Input().Get("pictureId"), 10, 64)
	isCorver, _ := strconv.ParseBool(this.Input().Get("isCorver"))

	err := models.IsCover(pictureId, productId, isCorver)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 设置封面成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

func (this *AdvertiseController) getCovers() {
	covers, err := models.GetCorver()
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "covers": covers, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *AdvertiseController) addAdvertise() {
	advertise := new(models.TAdvertise)
	this.ParseForm(advertise)
	advertiseId, err := models.AddAdvertise(advertise)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "advertiseId": advertiseId, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *AdvertiseController) delAdvertise() {
	advertiseId, _ := strconv.ParseInt(this.Input().Get("advertiseId"), 10, 64)
	err := models.DelAdvertise(advertiseId)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "删除广告成功!", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *AdvertiseController) mdfyAdvertise() {
	advertise := new(models.TAdvertise)
	this.ParseForm(advertise)
	num, err := models.MdfyAdvertise(advertise)
	beego.Info(num)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改广告成功!", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *AdvertiseController) getAdvertiseByPage() {
	index, _ := strconv.Atoi(this.Input().Get("index"))
	size, _ := strconv.Atoi(this.Input().Get("size"))
	where := this.Input().Get("where")

	advertises, count, err := models.GetAdvertiseByPage(index, size, where)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "advertises": advertises, "count": count, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}
