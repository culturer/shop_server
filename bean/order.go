package bean

import (
	"shop/models"
)

type OrderBean struct {
	Order      *models.TOrder
	Partner    *PartnerBean
	OrderItems []*OrderItemBean
}

func GetOrderBean(orderId int64) (*OrderBean, error) {

	order, err := models.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	mOrderItems, _, err := models.GetOrderItemByOrderId(orderId, 0, 100, "")
	if err != nil {
		return nil, err
	}

	orderItems := make([]*OrderItemBean, len(mOrderItems))
	for i := 0; i < len(mOrderItems); i++ {
		orderItemId := mOrderItems[i].Id
		orderItem, err := GetOrderItemBean(orderItemId)
		if err != nil {
			return nil, err
		}
		orderItems[i] = orderItem
	}

	partner, err := GetPartnerBean(order.PartnerId)
	if err != nil {
		return nil, err
	}
	orderBean := &OrderBean{Order: order, Partner: partner, OrderItems: orderItems}
	return orderBean, nil
}
