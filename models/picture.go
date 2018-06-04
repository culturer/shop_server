package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	// "fmt"
)

//产品图片
type TPicture struct {
	Id int64
	//产品Id
	ProductId int64
	//图片链接
	Url string
	//封面
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
	picture := &TPicture{Id: pictureId}
	_, err := o.Delete(picture)
	return err
}
func DelPictureByProductId(productId int64) error {
	o := orm.NewOrm()
	picture := &TPicture{ProductId: productId}
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

func GetPictureById(pictureId int64) (*TPicture, error) {
	o := orm.NewOrm()
	picture := new(TPicture)
	qs := o.QueryTable("t_picture")
	err := qs.Filter("id", pictureId).One(picture)
	return picture, err
}

//是否首页轮播
func IsCover(pictureId, productId int64, isCorver bool) error {
	picture, err := GetPictureById(pictureId)
	if err != nil {
		return err
	}
	if picture != nil {
		picture.IsCover = isCorver
	}
	o := orm.NewOrm()
	_, err = o.Update(picture)
	return err
}

func GetCorver() ([]*TPicture, error) {
	pictures := make([]*TPicture, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_picture")
	_, err := qs.Filter("is_cover", 1).All(&pictures)
	return pictures, err
}
