package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"path"
	"shop/models"
	"strconv"

	"time"
)

type PictureController struct {
	BaseController
}

func (this *PictureController) Get() {
	this.TplName = "login.html"
}

func (this *PictureController) Post() {

	options, _ := strconv.Atoi(this.Input().Get("options"))
	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	if options == 0 {
		productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
		pictures, err := this.getPicturesByProductId(productId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "查询图片异常，请检查参数后稍后再试！", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "pictures": pictures, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 1 {
		strProductId := this.Input().Get("productId")
		productId, _ := strconv.ParseInt(strProductId, 10, 64)
		isCover, _ := strconv.ParseBool(this.Input().Get("isCover"))
		//创建用户目录
		err := os.MkdirAll("pictures/"+strProductId, os.ModePerm)
		if err != nil {
			beego.Error(err)
		}
		// 获取附件
		_, fh, ee := this.GetFile("attachment")
		if ee != nil {
			beego.Error(ee)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": ee.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		var attachment string
		if fh != nil {
			//保存附件
			attachment = fh.Filename
			beego.Info(attachment)
			myPath := path.Join("pictures/"+strProductId, attachment)
			beego.Info(myPath)
			err := this.SaveToFile("attachment", myPath)

			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			pictureId, err := this.addPicture(productId, myPath, isCover)
			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "pictureId": pictureId, "url": myPath, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		pictureId, _ := strconv.ParseInt(this.Input().Get("pictureId"), 10, 64)
		err := this.delPicture(pictureId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "delete picture fail!", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "delete picture success!", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		_strProductId := this.Input().Get("productId")
		_productId, _ := strconv.ParseInt(_strProductId, 10, 64)
		_err := models.DelPictureByProductId(_productId)
		if _err != nil {
			beego.Error(_err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "删除原图片失败", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		strProductId := this.Input().Get("productId")
		productId, _ := strconv.ParseInt(strProductId, 10, 64)
		isCover, _ := strconv.ParseBool(this.Input().Get("isCover"))
		//创建用户目录
		err := os.MkdirAll("pictures/"+strProductId, os.ModePerm)
		if err != nil {
			beego.Error(err)
		}
		// 获取附件
		_, fh, ee := this.GetFile("attachment")
		if ee != nil {
			beego.Error(ee)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": ee.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		var attachment string
		if fh != nil {
			//保存附件
			attachment = fh.Filename
			beego.Info(attachment)
			myPath := path.Join("pictures/"+strProductId, attachment)
			beego.Info(myPath)
			err := this.SaveToFile("attachment", myPath)

			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			pictureId, err := this.addPicture(productId, myPath, isCover)
			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "pictureId": pictureId, "url": myPath, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

}

func (this *PictureController) addPicture(productId int64, url string, isCover bool) (int64, error) {
	pictureId, err := models.AddPicture(productId, url, isCover)
	return pictureId, err
}

func (this *PictureController) delPicture(pictureId int64) error {
	err := models.DelPicture(pictureId)
	return err
}

func (this *PictureController) getPicturesByProductId(productId int64) ([]*models.TPicture, error) {
	pictures, err := models.GetPicturesByProductId(productId)
	return pictures, err
}
