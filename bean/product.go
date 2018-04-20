package bean

import (
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
	productType, err := GetProductTypeBean(product.ProductTypeId)
	productBean := &ProductBean{Product: product, Pirctures: pirectures, ProductType: productType}
	return productBean, err
}
