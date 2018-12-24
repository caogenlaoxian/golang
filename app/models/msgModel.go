package models

import (
	. "../../app/database"
)

type Msg struct {
	Id      int
	Title   string
	Content string
	Start   int
	Limit   int
}

//查询列表
func (q Msg) GetMsgList() (lists []Msg, err error) {
	lists = make([]Msg, 0) //创建一个切片
	rows, err := SqlDB.Query("select * from msg ")
	if err != nil {
		return
	}

	for rows.Next() {
		var query Msg
		rows.Scan(&query.Id, &query.Title, &query.Content)
		lists = append(lists, query)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//带分页查询列表
func GetMsgList1(pageNum, pageSize int) (lists []Msg, err error) {
	lists = make([]Msg, 0) //创建一个切片
	rows, err := SqlDB.Query("select * from msg limit ?,?", (pageNum-1)*pageSize, pageSize)
	if err != nil {
		return
	}

	for rows.Next() {
		var query Msg
		rows.Scan(&query.Id, &query.Title, &query.Content)
		lists = append(lists, query)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//添加列表

//修改列表

//删除列表
