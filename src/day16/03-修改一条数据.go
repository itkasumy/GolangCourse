package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("数据库连接成功！")

	stmt, err := db.Prepare("update student set name = ? where id = ?" )
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	result, _ := stmt.Exec("zhaoliu", 1005)
	n1, _ := result.RowsAffected()
	n2, _ := result.LastInsertId()
	fmt.Println(n1, n2)
}
