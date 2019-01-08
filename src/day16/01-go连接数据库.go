package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(172.168.30.3:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败...", err.Error())
		return
	}
	fmt.Println("数据库连接成功！")
	defer db.Close()

	//fmt.Printf("%d", num)
	result, err := db.Exec("insert into student values (?, ?, ?, ?, ?)", 1003, "lisi", "19", 1, 611)
	if err != nil {
		fmt.Println("插入数据失败:", err.Error())
		return
	}

	i6401, err := result.LastInsertId()
	i6402, err := result.RowsAffected()
	fmt.Println(i6401, i6402)
}
