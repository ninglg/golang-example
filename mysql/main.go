package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error

	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True"
	// sql.Open仅创建DB对象，不会打开与数据库的任何连接，即不会校验账号密码是否正确。
	// 注意！这里不要使用:=，此处是给全局变量赋值，然后在main函数中使用全局变量db。
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
}

func main() {
	// 如果要测试连接，必须执行查询以强制打开连接。常见的方法是在数据库对象上调用Ping()。
	err := db.Ping()
	if err != nil {
		log.Panic("err:", err.Error())
	}

	// Delete

	// Insert

	// Query

	// Update

}