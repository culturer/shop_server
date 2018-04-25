package bean

import (
	"github.com/astaxie/beego"
	"shop/models"
)

type ProductBean struct {
	Product     *models.TProduct
	Pirctures   []*models.TPicture
	ProductType *ProductTypeBean
}

func GetProductBean(productId int64) (*ProductBean, error) {
	product, err := models.GetProductById(productId)

	pirectures, err := models.GetPicturesByProductId(productId)
	beego.Info(product.ProductTypeId)
	productType, err := GetProductTypeBean(product.ProductTypeId)
	beego.Info(productType)
	productBean := &ProductBean{Product: product, Pirctures: pirectures, ProductType: productType}
	return productBean, err
}
