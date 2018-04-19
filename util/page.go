package util

import ()

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

//分页获取数据
// func GetPage(user_id, size, index int, draft_type, where string) ([]interface{}, error) {
// //orm模块
// ormHelper := orm.NewOrm()
// //返回data数据
// data := []interface{}
// //错误对象
// _, err := ormHelper.Raw("select * from tb_draft where user_id=? and draft_type=? ? order by id desc limit ? offset ?", user_id, draft_type, where, size, size*(index-1)).QueryRows(&data)
// if err != nil {
// 	fmt.Printf("error is %v\n", err)
// }

// return data, err
// }
