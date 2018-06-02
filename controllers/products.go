package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"path"
	"shop/models"
	"strconv"
	"time"
)

type ProductController struct {
	BaseController
}

func (this *ProductController) Get() {
	var page string
	this.Ctx.Input.Bind(&page, "page")
	if page == "product_add" {
		this.TplName = "product_add.html"
	} else if page == "product_list" {
		this.TplName = "product_list.html"
	} else if page == "product_type_list" {
		this.TplName = "product_type_list.html"
	} else if page == "product_edit" {
		this.TplName = "product_edit.html"
	}
}

func (this *ProductController) Post() {

	// [types == 0  获取商品分类]
	// [types == 1  获取商品]
	types, _ := strconv.Atoi(this.Input().Get("types"))
	options, _ := strconv.Atoi(this.Input().Get("options"))

	// [types == 0  获取商品分类列表]
	if types == 0 {

		// [options == 0  查询]
		// [options == 1  增加]
		// [options == 2  删除]
		// [options == 3  修改]

		if options == 0 {

			//partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
			pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))

			productTypes, totalPage, err := this.getProductTypes(pageNo, pageSize, "")

			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品分类列表失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "productTypes": productTypes, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}

		if options == 1 {
			// sortId, _ := strconv.ParseInt(this.Input().Get("sortId"), 10, 64)
			// typeName := this.Input().Get("typeName")
			//创建用户目录
			err := os.MkdirAll("pictures/typeicon", os.ModePerm)
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
			var attachment, icon string
			if fh != nil {
				//保存附件
				attachment = fh.Filename
				beego.Info(attachment)
				myPath := path.Join("pictures/typeicon", attachment)
				beego.Info(myPath)
				err := this.SaveToFile("attachment", myPath)

				if err != nil {
					beego.Error(err)
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				icon = myPath

			}
			productType := new(models.TProductType)
			this.ParseForm(productType)
			productType.Icon = icon
			productTypeId, err := this.addProductType(productType)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 添加商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "productTypeId": productTypeId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			//this.Ctx.Output.JSON(data, hasIndent, coding)
			return
		}

		if options == 2 {
			productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			err := this.delProductType(productTypeId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 删除商品分类成功,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if options == 3 {
			productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			productType, _ := models.GetProductTypeById(productTypeId)
			// 获取附件
			_, fh, ee := this.GetFile("attachment")
			if ee != nil {
				beego.Error(ee)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": ee.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			var attachment, icon string
			if fh != nil {
				//保存附件
				attachment = fh.Filename
				beego.Info(attachment)
				myPath := path.Join("pictures/typeicon", attachment)
				beego.Info(myPath)
				err := this.SaveToFile("attachment", myPath)

				if err != nil {
					beego.Error(err)
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				icon = myPath

			}
			productType.Icon = icon
			this.ParseForm(productType)
			//util.formToModel(obj interface{},ctx *context)
			_, err := models.EditProductType(productType)
			beego.Info(productType)
			if err != nil {
				beego.Info(productType)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

			// mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
			// // [mdfyType == 0  修改分销商]
			// // [mdfyType == 1  修改优先级]
			// if mdfyType == 0 {
			// 	partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
			// 	productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			// 	err := this.mdfyProductTypePartner(productTypeId, partnerId)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类成功,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 1 {
			// 	sortId, _ := strconv.ParseInt(this.Input().Get("sortId"), 10, 64)
			// 	productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			// 	err := models.MdfyProductTypeSortId(productTypeId, sortId)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类优先级失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类优先级成功,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }
		}

	}

	// [types == 1  获取商品]
	if types == 1 {

		// [options == 0  查询]
		// [options == 1  增加]
		// [options == 2  删除]
		// [options == 3  修改]

		if options == 0 {

			getType, _ := strconv.Atoi(this.Input().Get("getType"))
			// [getType == 0 获取商品列表]
			// [getType == 1 获取指定商品]
			// [getType == 2 获取特色商品]
			// [getType == 3 获取特价商品]
			// [getType == 4 获取抢购商品]

			if getType == 0 {
				productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))
				where := this.Input().Get("where")

				products, totalPage, err := this.getProducts(productTypeId, pageNo, pageSize, where)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品列表失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				beego.Info(products)
				this.Data["json"] = map[string]interface{}{"status": 200, "products": products, "totalPage": totalPage, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if getType == 1 {
				productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
				product, err := this.getProduct(productId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}

				this.Data["json"] = map[string]interface{}{"status": 200, "product": product, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			if getType == 2 {
				// [getType == 2 获取特色商品]
				products, err := models.GetTeShe()
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				beego.Info(products)
				this.Data["json"] = map[string]interface{}{"status": 200, "products": products, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			if getType == 3 {
				// [getType == 3 获取特价商品]
				products, err := models.GetTejia()
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				beego.Info(products)
				this.Data["json"] = map[string]interface{}{"status": 200, "products": products, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			if getType == 4 {
				// [getType == 4 获取抢购商品]
				products, err := models.GetQiangGou()
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				beego.Info(products)
				this.Data["json"] = map[string]interface{}{"status": 200, "products": products, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		}

		if options == 1 {
			product := new(models.TProduct)
			this.ParseForm(product)
			product.UserId, _ = this.GetSession("uid").(int64)
			productId, err := this.addProduct(product)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 新增商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "productId": productId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if options == 2 {
			// [delType == 0  删除产品]
			// [delType == 1  删除产品图片]
			delType, _ := strconv.Atoi(this.Input().Get("delType"))

			if delType == 0 {

				productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)
				err := this.delProduct(productId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}

				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 删除商品成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return

			}

			if delType == 1 {
				pictureId, _ := strconv.ParseInt(this.Input().Get("pictureId"), 10, 64)
				err := models.DelPicture(pictureId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 删除图片失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 删除图片成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		}

		if options == 3 {
			//mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
			// [mdfyType == 0  修改分类]
			// [mdfyType == 1  修改名称]
			// [mdfyType == 2  修改库存]
			// [mdfyType == 3  修改价格]
			// [mdfyType == 4  修改成本]
			// [mdfyType == 5  修改描述]
			// [mdfyType == 6  修改备注]
			// [mdfyType == 7  修改优先级]
			productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)

			product, _ := models.GetProductById(productId)

			this.ParseForm(product)
			//util.formToModel(obj interface{},ctx *context)
			//beego.Info(product)
			_, err := models.EditProduct(product)

			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "product_id": productId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

			// if mdfyType == 0 {
			// 	productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			// 	err := models.MdfyType(productId, productTypeId)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 1 {
			// 	name := this.Input().Get("name")
			// 	err := models.MdfyName(productId, name)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品名称失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品名称成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 2 {
			// 	count, _ := strconv.Atoi(this.Input().Get("count"))
			// 	err := models.MdfyCount(productId, count)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品库存失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品库存成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 3 {
			// 	price, err := strconv.ParseFloat(this.Input().Get("price"), 32)
			// 	err = models.MdfyPrice(productId, price)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品价格失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品价格成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 4 {
			// 	standardPrice, err := strconv.ParseFloat(this.Input().Get("standardPrice"), 32)
			// 	err = models.MdfyStandardPrice(productId, standardPrice)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品成本失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品成本成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 5 {
			// 	desc := this.Input().Get("desc")
			// 	err := models.MdfyDesc(productId, desc)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品描述失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品描述成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 6 {
			// 	msg := this.Input().Get("msg")
			// 	err := models.MdfyMsg(productId, msg)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品备注失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品备注成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }

			// if mdfyType == 7 {
			// 	sortId, _ := strconv.Atoi(this.Input().Get("sortId"))
			// 	err := models.MdfyProductSort(productId, sortId)
			// 	if err != nil {
			// 		beego.Info(err.Error())
			// 		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品优先级失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 		this.ServeJSON()
			// 		return
			// 	}
			// 	this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品优先级成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
			// 	this.ServeJSON()
			// 	return
			// }
		}

	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 参数错误，请检查后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *ProductController) getProductTypes(pageNo, pageSize int, where string) ([]*models.TProductType, int, error) {
	productTypes, totalPage, err := models.GetProductTypePage(pageNo, pageSize, where)
	return productTypes, totalPage, err
}

func (this *ProductController) addProductType(productType *models.TProductType) (int64, error) {
	productTypeId, err := models.AddProductType(productType)
	return productTypeId, err
}

func (this *ProductController) delProductType(productTypeId int64) error {
	err := models.DelProductType(productTypeId)
	return err
}

func (this *ProductController) mdfyProductTypePartner(productTypeId int64, sortId int64) error {
	err := models.MdfyPartner(productTypeId, sortId)
	return err
}

func (this *ProductController) getProduct(productId int64) (*models.TProduct, error) {
	product, err := models.GetProductById(productId)
	return product, err
}

func (this *ProductController) getProducts(productTypeId int64, pageNo, pageSize int, where string) ([]*models.TProduct, int, error) {
	products, totalPage, err := models.GetProductByType(productTypeId, pageNo, pageSize, where)
	return products, totalPage, err
}

func (this *ProductController) addProduct(product *models.TProduct) (int64, error) {
	productId, err := models.AddProduct(product)
	return productId, err
}

func (this *ProductController) delProduct(productId int64) error {
	err := models.DelProduct(productId)
	return err
}
