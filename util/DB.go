package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

//InitDB 初始化数据库
func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/GinWebPhoto?charset=utf8")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(12)
	Db.SetMaxIdleConns(3)
	err = Db.Ping()
	if err != nil {
		log.Panicln("数据库链接出错！ error: ", err)
	}
}
