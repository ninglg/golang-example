package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//type student struct {
//	stuId int ``
//	stuName string 	``
//}

func init() {
	var err error

	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True"
	// sql.Open仅创建DB对象，不会打开与数据库的任何连接，即不会校验账号密码是否正确。
	// 注意！这里不要使用:=，此处是给全局变量赋值，然后在main函数中使用全局变量db。
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panicln("error:", err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
}

func main() {
	var err error

	defer db.Close()
	// 如果要测试连接，必须执行查询以强制打开连接。常见的方法是在数据库对象上调用Ping()。
	err = db.Ping()
	if err != nil {
		fmt.Println("db.Ping error:", err.Error())
		return
	}

	fmt.Println("db.Ping success!")

	// 0.使用db.Exec方法来进行Delete
	result, err := db.Exec("delete from student")
	if err != nil {
		fmt.Println("db.Exec delete error:", err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("db.Exec delete rowsAffected is:", rowsAffected)

	// 1.使用db.Exec方法来直接进行Insert
	result, err = db.Exec("insert student (stu_id, stu_name) values (?, ?)", 100, "xiaogang")
	if err != nil {
		fmt.Println("db.Exec insert error:", err.Error())
		return
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("db.Exec insert lastId is: ", lastId)

	// 2.使用Prepare方法和stmt的Exec方法来进行Insert
	stmt, err := db.Prepare("insert student (stu_id, stu_name) values (?, ?)")
	if err != nil {
		fmt.Println("db.Prepare error:", err.Error())
		return
	}
	// 资源关闭
	defer stmt.Close()

	result, _ = stmt.Exec(200, "xiaoming")
	lastId, _ = result.LastInsertId()
	fmt.Println("db.Prepare stmt.Exec insert lastId is: ", lastId)

	// 3.使用db.Exec方法来进行Update
	result, err = db.Exec("update student set stu_name='xiaoqiang' where stu_id = ?", 200)
	if err != nil {
		fmt.Println("db.Exec update error:", err.Error())
		return
	}

	rowsAffected, _ = result.RowsAffected()
	fmt.Println("db.Exec update rowsAffected is:", rowsAffected)

	// 4.使用db.QueryRow查询一条记录
	var stuId int
	var stuName string
	db.QueryRow("select stu_id, stu_name from student where stu_id = ?", 100).Scan(&stuId, &stuName)
	fmt.Println("db.QueryRow:", stuId, stuName)

	// 5.使用db.Query查询多条记录
	rows, err := db.Query("select stu_id, stu_name from student")
	if err != nil {
		fmt.Println("db.Query rows error:", err.Error())
		return
	}

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

}