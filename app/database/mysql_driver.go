package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

//初始化
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "beta-mysql:nopass.2@tcp(10.255.72.27:3306)/test?charset=utf8")
	//SqlDB, err = sql.Open("mysql", "root:dashayu@tcp(127.0.0.1:3306)/test?charset=utf8")
	// SqlDB, err = sql.Open("mysql", "beta-mysql:nopass.2@tcp(10.255.72.27:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
