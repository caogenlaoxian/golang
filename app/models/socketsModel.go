package models

import (
	. "../../app/database"
)
type Sockets struct{
	Name string
	Code string
}

type SocketsDetail struct{
	SocketCode string
	Kai string
	Shou string
	Height string
	Low string
	Liang string
	Jia string
	SocketDate string
}

//查询总数
func SocketsTotal() (total int64){
	rows   := SqlDB.QueryRow("SELECT count(*) as total from socket")
	rows.Scan(&total)
	return
}

//查询
func SocketList(pageNum,pageSize int) (lists []Sockets, err error){
	lists = make([]Sockets, 0) //创建一个切片
	offset := pageNum * pageSize
	rows ,_ := SqlDB.Query("SELECT name as Name,code as Code from socket limit ?,?",offset,pageSize)
	for rows.Next(){
		var query Sockets
		rows.Scan(&query.Name,&query.Code)
		lists = append(lists, query)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//truncate
func (s *Sockets) SocketTruncate(){
	SqlDB.Exec("TRUNCATE socket")
}

//add
func (s *Sockets) SocketsAdd() (lastId int64,err error){
	rs, err := SqlDB.Exec("INSERT INTO socket(name,code)VALUES(?,?)", s.Name, s.Code)
	if err != nil {
		return
	}
	lastId, err = rs.LastInsertId()
	return lastId, err
}

func (s *SocketsDetail) SocketDetailAdd() (lastId int64,err error){
	rs ,err := SqlDB.Exec("INSERT INTO socket_detail(socket_code,kai,shou,height,low,liang,jia,socket_date)VALUES(?,?,?,?,?,?,?,?)", s.SocketCode, s.Kai,s.Shou,s.Height,s.Low,s.Liang,s.Jia,s.SocketDate)
	if err != nil{
		return
	}
	lastId , err = rs.LastInsertId()
	return lastId , err
}
