package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//产品图片
type TPicture struct {
	Id int64
	//产品Id
	ProductId int64
	//图片链接
	Url string
	//封面?
	IsCover bool
}

func AddPicture(productId int64, url string, isCover bool) (int64, error) {
	o := orm.NewOrm()
	picture := &TPicture{ProductId: productId, Url: url, IsCover: isCover}
	pictureId, err := o.Insert(picture)
	return pictureId, err
}

func DelPicture(pictureId int64) error {
	o := orm.NewOrm()
	picture := &TPartner{Id: pictureId}
	_, err := o.Delete(picture)
	return err
}

func GetPicturesByProductId(productId int64) ([]*TPicture, error) {
	pictures := make([]*TPicture, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_picture")
	_, err := qs.Filter("product_id", productId).All(&pictures)
	return pictures, err
}

func MdfyPictureIsCover(productId int64, isCover bool) error {
	return nil
}
