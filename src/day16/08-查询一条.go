package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
		return
	}
	defer db.Close()

	row := db.QueryRow("select id, name, age, score from student where id = ?", 1005)
	var id, age, score int
	var name string

	err = row.Scan(&id, &name, &age, &score)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, name, age, score)
}
