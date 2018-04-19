package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"strconv"

	"time"
)

type AddressController struct {
	BaseController
}

func (this *AddressController) Get() {
	this.TplName = "login.html"
}

func (this *AddressController) Post() {

	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	options, _ := strconv.Atoi(this.Input().Get("options"))

	if options == 0 {
		// [getType == 0  根据addressId查询]
		// [getType == 1  根据userId查询]
		getType, _ := strconv.Atoi(this.Input().Get("getType"))
		if getType == 0 {
			addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
			address, err := this.getAddressById(addressId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询地址失败，请稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "address": address, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		if getType == 1 {

			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
			pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))

			addresses, totalPage, err := this.getAddressByUserId(userId, pageNo, pageSize, "")
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 查询地址失败，请稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "addresses": addresses, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}

	if options == 1 {
		userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
		country := this.Input().Get("country")
		province := this.Input().Get("province")
		city := this.Input().Get("city")
		block := this.Input().Get("block")
		street := this.Input().Get("street")
		community := this.Input().Get("community")
		desc := this.Input().Get("desc")
		tel := this.Input().Get("tel")
		name := this.Input().Get("name")
		addressId, err := this.addAddress(userId, country, province, city, block, street, community, desc, tel, name)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 新增地址失败，请稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "addressId": addressId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		err := this.delAddress(addressId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除地址失败，请稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "删除地址成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 3 {
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		sortId, _ := strconv.Atoi(this.Input().Get("sortId"))
		err := this.mdfyAddressSort(addressId, sortId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "修改地址优先级失败，请稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "修改地址优先级成功！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

}

func (this *AddressController) addAddress(userId int64, country string, province string, city string, block string, street string, community string, desc string, tel string, name string) (int64, error) {
	addressId, err := models.AddAddress(userId, country, province, city, block, street, community, desc, tel, name)
	return addressId, err
}

func (this *AddressController) delAddress(addressId int64) error {
	err := models.DelAddress(addressId)
	return err
}

func (this *AddressController) getAddressById(addressId int64) (*models.TAddress, error) {
	address, err := models.GetAddressById(addressId)
	return address, err
}

func (this *AddressController) getAddressByUserId(userId int64, pageNo, pageSize int, where string) ([]*models.TAddress, int, error) {
	address, totalPage, err := models.GetAddressByUserId(userId, pageNo, pageSize, "")
	return address, totalPage, err
}

func (this *AddressController) mdfyAddressSort(addressId int64, sortId int) error {
	err := models.MdfyAddressSort(addressId, sortId)
	return err
}
