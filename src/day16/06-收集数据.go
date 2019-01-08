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

	rows, err := db.Query("select id, name, age, score from student where age = ?", 18)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
		return
	}
	defer rows.Close()

	fields, _ := rows.Columns()
	for i := 0; i < len(fields); i++ {
		fmt.Print(fields[i], "\t\t")
	}
	fmt.Println()
	slice := make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var name string
		var age int
		var score int
		mp := make(map[string]interface{})

		err := rows.Scan(&id, &name, &age, &score)
		if err != nil {
			fmt.Println(err)
		}
		mp["id"] = id
		mp["name"] = name
		mp["age"] = age
		mp["score"] = score

		slice = append(slice, mp)
	}

	for _, val := range slice {
		for _, v := range val{
			fmt.Print(v, "\t\t")
		}
		fmt.Println()
	}
}
