package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type stuInfo struct {
	id int
	name string
	age int
	score int
}

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, age, score from student where age = ?", 18)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
		return
	}
	defer rows.Close()

	fields, _ := rows.Columns()
	for i := 0; i < len(fields); i++ {
		fmt.Print(fields[i], "\t\t\t")
	}
	fmt.Println()
	slice := make([]stuInfo, 0)
	for rows.Next() {
		var id int
		var name string
		var age int
		var score int

		err := rows.Scan(&id, &name, &age, &score)
		if err != nil {
			fmt.Println(err)
		}
		oneRow := stuInfo{id, name, age, score}
		//fmt.Println(oneRow)
		slice = append(slice, oneRow)
	}

	for _, val := range slice {
		fmt.Print(val.id, "\t\t\t")
		fmt.Print(val.name, "\t\t\t")
		fmt.Print(val.age, "\t\t\t")
		fmt.Print(val.score, "\t\t\t")
		fmt.Println()
	}
}
