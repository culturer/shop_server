package controllers

import (
	"github.com/astaxie/beego"
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
			sortId, _ := strconv.ParseInt(this.Input().Get("sortId"), 10, 64)
			typeName := this.Input().Get("typeName")
			productTypeId, err := this.addProductType(typeName, sortId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 添加商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "productTypeId": productTypeId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
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
			mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
			// [mdfyType == 0  修改分销商]
			// [mdfyType == 1  修改优先级]
			if mdfyType == 0 {
				partnerId, _ := strconv.ParseInt(this.Input().Get("partnerId"), 10, 64)
				productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
				err := this.mdfyProductTypePartner(productTypeId, partnerId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类成功,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 1 {
				sortId, _ := strconv.ParseInt(this.Input().Get("sortId"), 10, 64)
				productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
				err := models.MdfyProductTypeSortId(productTypeId, sortId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类优先级失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类优先级成功,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
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

			if getType == 0 {
				productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
				pageNo, _ := strconv.Atoi(this.Input().Get("pageNo"))
				pageSize, _ := strconv.Atoi(this.Input().Get("pageSize"))

				products, totalPage, err := this.getProducts(productTypeId, pageNo, pageSize)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 获取商品列表失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}

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

		}

		if options == 1 {
			productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
			userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			name := this.Input().Get("name")
			price, _ := strconv.ParseFloat(this.Input().Get("price"), 32)
			desc := this.Input().Get("desc")
			count, _ := strconv.Atoi(this.Input().Get("count"))
			standardPrice, _ := strconv.ParseFloat(this.Input().Get("standardPrice"), 32)
			msg := this.Input().Get("msg")
			productId, err := this.addProduct(productTypeId, userId, name, count, standardPrice, price, desc, msg)
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

		if options == 3 {
			mdfyType, _ := strconv.Atoi(this.Input().Get("mdfyType"))
			// [mdfyType == 0  修改分类]
			// [mdfyType == 1  修改名称]
			// [mdfyType == 2  修改库存]
			// [mdfyType == 3  修改价格]
			// [mdfyType == 4  修改成本]
			// [mdfyType == 5  修改描述]
			// [mdfyType == 6  修改备注]
			// [mdfyType == 7  修改优先级]
			productId, _ := strconv.ParseInt(this.Input().Get("productId"), 10, 64)

			if mdfyType == 0 {
				productTypeId, _ := strconv.ParseInt(this.Input().Get("productTypeId"), 10, 64)
				err := models.MdfyType(productId, productTypeId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品分类失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品分类成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 1 {
				name := this.Input().Get("name")
				err := models.MdfyName(productId, name)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品名称失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品名称成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 2 {
				count, _ := strconv.Atoi(this.Input().Get("count"))
				err := models.MdfyCount(productId, count)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品库存失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品库存成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 3 {
				price, err := strconv.ParseFloat(this.Input().Get("price"), 32)
				err = models.MdfyPrice(productId, price)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品价格失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品价格成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 4 {
				standardPrice, err := strconv.ParseFloat(this.Input().Get("standardPrice"), 32)
				err = models.MdfyStandardPrice(productId, standardPrice)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品成本失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品成本成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 5 {
				desc := this.Input().Get("desc")
				err := models.MdfyDesc(productId, desc)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品描述失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品描述成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 6 {
				msg := this.Input().Get("msg")
				err := models.MdfyMsg(productId, msg)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品备注失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品备注成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			if mdfyType == 7 {
				sortId, _ := strconv.Atoi(this.Input().Get("sortId"))
				err := models.MdfyProductSort(productId, sortId)
				if err != nil {
					beego.Info(err.Error())
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 修改商品优先级失败,请稍后再试！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
					this.ServeJSON()
					return
				}
				this.Data["json"] = map[string]interface{}{"status": 200, "msg": " 修改商品优先级成功！ ", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
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

func (this *ProductController) addProductType(typeName string, sortId int64) (int64, error) {
	productTypeId, err := models.AddProductType(typeName, sortId)
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

func (this *ProductController) getProducts(productTypeId int64, pageNo, pageSize int) ([]*models.TProduct, int, error) {
	products, totalPage, err := models.GetProductByType(productTypeId, pageNo, pageSize, "")
	return products, totalPage, err
}

func (this *ProductController) addProduct(productTypeId int64, userId int64, name string, count int, standardPrice float64, price float64, desc string, msg string) (int64, error) {
	productId, err := models.AddProduct(productTypeId, userId, name, count, standardPrice, price, desc, msg)
	return productId, err
}

func (this *ProductController) delProduct(productId int64) error {
	err := models.DelProduct(productId)
	return err
}
