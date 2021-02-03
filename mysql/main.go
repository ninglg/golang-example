package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// sql.DB的设计就是用来作为长连接使用的，不要频繁Open和Close。
// 比较好的做法是，为每个不同的datastore建一个DB对象，保持这些对象Open。
// 如果需要短连接，那么把DB作为参数传入function，而不要在function中Open和Close。
var db *sql.DB

func init() {
	// Driver Name
	driverName := "mysql"
	// DSN: Data Source Name
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True"
	// sql.Open仅创建DB对象，不会打开与数据库的任何连接，即不会校验账号密码是否正确。
	// 注意！这里不要使用:=，此处是给全局变量赋值，然后在main函数中使用全局变量db。
	db, _ = sql.Open(driverName, dataSourceName)

	// 设置数据库最大连接数
	db.SetMaxOpenConns(10)
	// 设置数据库最大空闲连接数
	db.SetMaxIdleConns(10)
}

func ping() {
	// 如果要测试连接，必须执行查询以强制打开连接。常见的方法是在数据库对象上调用Ping()。
	err := db.Ping()
	if err != nil {
		fmt.Println("db.Ping error:", err.Error())
		return
	}

	fmt.Println("db.Ping success!")
}

func deleteSql() {
	// 使用db.Exec方法来进行Delete
	result, err := db.Exec("delete from student")
	if err != nil {
		fmt.Println("db.Exec delete error:", err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("db.Exec delete rowsAffected is:", rowsAffected)
}

func insertSql() {
	// 使用db.Exec方法来直接进行Insert
	result, err := db.Exec("insert student (stu_id, stu_name) values (?, ?)", 100, "xiaogang")
	if err != nil {
		fmt.Println("db.Exec insert error:", err.Error())
		return
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("db.Exec insert lastId is: ", lastId)
}

func insertSql2() {
	// 一般用Prepared Statements和stmt.Exec()来完成insert, update, delete操作。
	// Prepared Statements是和单个数据库连接绑定的。客户端发送一个有占位符的statement到服务端，服务器返回一个statement ID，然后客户端发送ID和参数来执行statement。
	// database/sql包有自动重试等功能。当生成一个Prepared Statement时会自动在连接池中绑定到一个空闲连接，Stmt对象记住了绑定的连接。当执行Stmt时会尝试使用该连接。如果不可用，例如连接被关闭或繁忙中，会自动re-prepare，绑定到另一个连接。
	stmt, err := db.Prepare("insert student (stu_id, stu_name) values (?, ?)")
	if err != nil {
		fmt.Println("db.Prepare error:", err.Error())
		return
	}
	defer stmt.Close()

	result, _ := stmt.Exec(200, "xiaoming")
	lastId, _ := result.LastInsertId()
	fmt.Println("db.Prepare stmt.Exec insert lastId is: ", lastId)
}

func updateSql() {
	// 使用db.Exec方法来进行Update
	result, err := db.Exec("update student set stu_name='xiaoqiang' where stu_id = ?", 200)
	if err != nil {
		fmt.Println("db.Exec update error:", err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("db.Exec update rowsAffected is:", rowsAffected)
}

func queryRow() {
	// 使用db.QueryRow查询一条记录
	var stuId int
	var stuName string
	_ = db.QueryRow("select stu_id, stu_name from student where stu_id = ?", 100).Scan(&stuId, &stuName)
	fmt.Println("db.QueryRow:", stuId, stuName)
}

func query() {
	// 使用db.Query查询多条记录
	rows, err := db.Query("select stu_id, stu_name from student")
	if err != nil {
		fmt.Println("db.Query rows error:", err.Error())
		return
	}
	// rows.Close()也可以多次调用，无副作用
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			name string
		)
		if err = rows.Scan(&id, &name); err != nil {
			fmt.Println("db.Query rows.Next error:", err.Error())
			continue
		}

		fmt.Println("db.Query rows.Next:", id, name)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("db.Query rows.Err:", err.Error())
	}
}

func trans() {
	// db.Begin()开始事务，Commit()或Rollback()关闭事务。
	// Tx从连接池中取出一个连接，在关闭之前都是使用这个连接。
	// Tx不能和DB层的Begin, Commit混合使用。

}

func main() {
	defer db.Close()

	ping()
	deleteSql()
	insertSql()
	insertSql2()
	updateSql()
	queryRow()
	query()
	trans()
}
