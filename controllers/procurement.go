package controllers

import (
	"github.com/astaxie/beego"
	// "shop/models"
	// "time"
)

type ProcurementController struct {
	beego.Controller
}

func (this *ProcurementController) Get() {
	this.TplName = "procurement.html"
	var page string
	this.Ctx.Input.Bind(&page, "page")
}
