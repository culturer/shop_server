package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

//广告
type TComment struct {
	Id int64
	//发表评论的用户
	UserId int64
	//回复的类型 adv --- 对广告的评论 advc --- 对广告的评论的评论
	Type string
	//回复的广告的Id
	AdvertiseId int64
	//回复的评论对象
	CommentId int64
	//关联的产品Id，方便直接打开产品信息
	ProductId int64
	//内容
	Content string
	//评论时间
	CreateTime string
}

func AddComment(comment *TComment) (int64, error) {
	//防止误设置id影响排序
	comment.Id = 0
	comment.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	commentId, err := o.Insert(comment)
	return commentId, err
}

func DelComment(commentId int64) error {
	o := orm.NewOrm()
	comment := &TComment{Id: commentId}
	_, err := o.Delete(comment)
	return err
}

func GetComments(index, size int, where string) ([]*TComment, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TComment{}
	dataCounts := []*TComment{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_comment %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_comment  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	return data, int(count), err
}
