package bean

import (
	"shop/models"
)

type UserBean struct {
	User    *models.TUser
	Address []*models.TAddress
	Orders  []*OrderBean
}

func GetUserBean(userId int64) (*UserBean, error) {
	mUser, err := models.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	address, _, err := models.GetAddressByUserId(userId, 100, 100, "")
	if err != nil {
		return nil, err
	}
	mOrders, _, err := models.GetOrderByUserId(userId, 100, 100, "")
	orders := make([]*OrderBean, len(mOrders))
	for i := 0; i < len(mOrders); i++ {
		orders[i], err = GetOrderBean(mOrders[i].Id)
		if err != nil {
			return nil, err
		}
	}
	user := &UserBean{User: mUser, Address: address, Orders: orders}
	return user, nil
}