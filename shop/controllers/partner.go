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
	this.TplName = "login.html"
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
			partners, err := this.getPartners()
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询分销商异常，请检查后重新查询！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "partners": partners, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	if options == 1 {
		userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		partnerName := this.Input().Get("partnerName")
		partnerId, err := this.addPartner(userId, partnerName, addressId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 添加分销商异常，请检查后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
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
		// [mdfyType == 3  修改分销商信用]
		// [mdfyType == 4  修改分销商权限]
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
			addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
			err := this.mdfyPartnerAddress(partnerId, addressId)
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
	}

}

func (this *PartnerController) addPartner(userId int64, partnerName string, addressId int64) (int64, error) {
	partnerId, err := models.AddPartner(userId, partnerName, addressId)
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

func (this *PartnerController) getPartners() ([]*models.TPartner, error) {
	partners, err := models.GetPartners()
	return partners, err
}

func (this *PartnerController) mdfyPartnerName(partnerId int64, partnerName string) error {
	err := models.MdfyPartnerName(partnerId, partnerName)
	return err
}

func (this *PartnerController) mdfyPartnerAddress(partnerId int64, addressId int64) error {
	err := models.MdfyPartnerAddress(partnerId, addressId)
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
