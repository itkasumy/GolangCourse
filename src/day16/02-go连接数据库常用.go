package main

// step01: 导包
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	//step02: 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(172.168.30.3:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败：", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("数据库连接成功！")

	//step03: 准备
	stmt, err := db.Prepare("insert into student (id, name, age, gender, score) values (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	//step04: 插入数据
	result, err := stmt.Exec(1005, "zhangliu", 17, 0, 521)
	if err != nil {
		fmt.Println(err)
	}
	result, err = stmt.Exec(1006, "chenqi", 18, 1, 500)

	n01, _ := result.LastInsertId()
	n02, _ := result.RowsAffected()
	//n03, _ := result2.LastInsertId()
	fmt.Println(n01, n02)
	if n02 == 1 {
		fmt.Println("插入数据成功!")
	}
}
