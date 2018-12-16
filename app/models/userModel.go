package models

import (
	. "../../app/database"
)

type User struct {
	Id        int    `json:"id" form "id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form: "last_name"`
}

//查询
func (u *User) GetUserInfo() (users []User, err error) {
	//创建一个切片
	users = make([]User, 0)
	rows, err := SqlDB.Query("select id, first_name, last_name from user")
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

//第二种查询方法
func (u *User) GetUserInfoById() (err error) {
	SqlDB.QueryRow("select id,first_name,last_name from user where id = ?", u.Id).Scan(&u.Id, &u.FirstName, &u.LastName)

	return
}

//插入
func (u *User) AddUser() (lastId int64, err error) {
	rs, err := SqlDB.Exec("INSERT INTO user(first_name,last_name)VALUES(?,?)", u.FirstName, u.LastName)
	if err != nil {
		return
	}
	lastId, err = rs.LastInsertId()
	return lastId, err
}

//修改
func (u *User) UpdateUser() (effect_num int64, err error) {
	rs, err := SqlDB.Exec("UPDATE user SET first_name= ?,last_name = ? where id = ?", u.FirstName, u.LastName)
	if err != nil {
		return
	}
	effect_num, err = rs.RowsAffected()
	return effect_num, err
}

//删除
func (u *User) DelUser() (effort_num int64, err error) {
	rs, err := SqlDB.Exec("DELETE FROM user where id= ?", u.Id)
	if err != nil {
		return
	}
	effort_num, err = rs.RowsAffected()
	return
}
