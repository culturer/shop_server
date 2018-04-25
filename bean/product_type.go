package bean

import (
	"shop/models"
)

type ProductTypeBean struct {
	ProductType  *models.TProductType
	ProductBeans []*ProductBean
	//Partner      *PartnerBean
}

func GetProductTypeBean(productTypeId int64) (*ProductTypeBean, error) {
	productType, err := models.GetProductTypeById(productTypeId)
	if err != nil {
		return nil, err
	}
	mProducts, _, err := models.GetProductByType(productTypeId, 100, 100, "")
	products := make([]*ProductBean, len(mProducts))
	for i := 0; i < len(mProducts); i++ {
		products[i], err = GetProductBean(mProducts[i].Id)
		if err != nil {
			return nil, err
		}
	}
	// partner, err := GetPartnerBean(productType.SortId)
	// if err != nil {
	// 	return nil, err
	// }
	productTypeBean := &ProductTypeBean{ProductType: productType, ProductBeans: products}
	return productTypeBean, err
}
