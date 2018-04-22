package bean

import (
	"shop/models"
)

type PartnerBean struct {
	Partner *models.TPartner
	User    *UserBean
	//Address *models.TAddress
}

func GetPartnerBean(partnerId int64) (*PartnerBean, error) {

	partenr, err := models.GetPartnerById(partnerId)
	user, err := GetUserBean(partenr.UserId)
	//address, err := models.GetAddressById(partenr.Address)
	partenrBean := &PartnerBean{Partner: partenr, User: user}
	return partenrBean, err

}
