package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"
	"time"
)

type PartnerController struct {
	BaseController
}

func (this *PartnerController) Get() {
	var page string
	this.Ctx.Input.Bind(&page, "page")
	if page == "partner_add" {
		this.TplName = "partner_add.html"
	} else if page == "partner_list" {
		this.TplName = "partner_list.html"
	}
}

func (this *PartnerController) Post() {

	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]

	options, _ := strconv.Atoi(this.Input().Get("options"))

	if options == 0 {
		// [getType == 0  根据partnerId查询]
		// [getType == 1  查询所有分销商]
		getType, _ := strconv.Atoi(this.Input().Get("getType"))
		if getType == 0 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			partner, err := this.getPartnerById(partnerId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询分销商异常，请检查后重新查询！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "partner": partner, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if getType == 1 {

			pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
			pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))

			partners, totalPage, err := this.getPartners(pageNo, pageSize)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询分销商异常，请检查后重新查询！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "partners": partners, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	if options == 1 {
		var userId int64
		tel := this.Input().Get("tel")
		user, err := models.GetUserByTel(tel)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		userId = user.Id
		// userId, _ = (this.GetSession("uid")).(int64)
		address := this.Input().Get("address")
		partnerName := this.Input().Get("partnerName")
		position := this.Input().Get("position")
		desc := this.Input().Get("desc")
		beego.Info(userId)
		partnerId, err := this.addPartner(userId, partnerName, address, position, desc)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "partnerId": partnerId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
		err := this.delPartner(partnerId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除分销商异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "删除分销商成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	if options == 3 {
		// [mdfyType == 0  修改分销商名称]
		// [mdfyType == 1  修改分销商地址]
		// [mdfyType == 2  修改分销商信用]
		// [mdfyType == 3  修改分销商权限]
		// [mdfyType == 4  修改分销商排序优先级]
		mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))

		if mdfyType == 0 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			partnerName := this.Input().Get("partnerName")
			err := this.mdfyPartnerName(partnerId, partnerName)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改分销商名称异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改分销商名称成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 1 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			address := this.Input().Get("address")
			err := this.mdfyPartnerAddress(partnerId, address)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改分销商地址异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改分销商地址成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 2 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			credits, _ := strconv.Atoi(this.Input().Get("credits"))
			err := this.mdfyPartnerCredits(partnerId, credits)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改分销商信用异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改分销商信用成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if mdfyType == 3 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			pro, _ := strconv.Atoi(this.Input().Get("pro"))
			err := this.mdfyPartnerPro(partnerId, pro)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改分销商权限异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改分销商权限成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		if mdfyType == 4 {
			partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			sortId, _ := strconv.Atoi(this.Input().Get("sortId"))
			err := this.mdfyPartnerSort(partnerId, sortId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改分销商排序优先级异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改分销商排序优先级成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}

}

func (this *PartnerController) addPartner(userId int64, partnerName string, address, position, desc string) (int64, error) {
	partnerId, err := models.AddPartner(userId, partnerName, address, position, desc)
	return partnerId, err
}

func (this *PartnerController) delPartner(partnerId int64) error {
	err := models.DelPartner(partnerId)
	return err
}

func (this *PartnerController) getPartnerById(partnerId int64) (*models.TPartner, error) {
	partner, err := models.GetPartnerById(partnerId)
	return partner, err
}

func (this *PartnerController) getPartners(pageNo, pageSize int) ([]*models.TPartner, int, error) {
	partners, totalPage, err := models.GetPartners(pageNo, pageSize)
	return partners, totalPage, err
}

func (this *PartnerController) mdfyPartnerName(partnerId int64, partnerName string) error {
	err := models.MdfyPartnerName(partnerId, partnerName)
	return err
}

func (this *PartnerController) mdfyPartnerAddress(partnerId int64, address string) error {
	err := models.MdfyPartnerAddress(partnerId, address)
	return err
}

func (this *PartnerController) mdfyPartnerCredits(partnerId int64, credits int) error {
	err := models.MdfyPartnerCredits(partnerId, credits)
	return err
}

func (this *PartnerController) mdfyPartnerPro(partnerId int64, pro int) error {
	err := models.MdfyPartnerPro(partnerId, pro)
	return err
}

func (this *PartnerController) mdfyPartnerSort(partnerId int64, sortId int) error {
	err := models.MdfyPartnerSort(partnerId, sortId)
	return err
}
