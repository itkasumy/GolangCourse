package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("数据库连接成功！")

	stmt, err := db.Prepare("delete from student where age = ?" )
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	result, _ := stmt.Exec(19)

	n1, _ := result.LastInsertId()
	n2, _ := result.RowsAffected()

	fmt.Println(n1)
	if n2 > 0 {
		fmt.Println("删除数据成功：", n2)
	}
}
