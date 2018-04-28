package bean

import (
	"github.com/astaxie/beego"
	"shop/models"
)

type ProductBean struct {
	Product     *models.TProduct
	Pictures    []*models.TPicture
	ProductType *models.TProductType
}

func GetProductBean(productId int64) (*ProductBean, error) {
	product, err := models.GetProductById(productId)

	pirectures, err := models.GetPicturesByProductId(productId)
	beego.Info(product.ProductTypeId)
	productType, err := models.GetProductTypeById(product.ProductTypeId)
	beego.Info(productType)
	productBean := &ProductBean{Product: product, Pictures: pirectures, ProductType: productType}
	return productBean, err
}
