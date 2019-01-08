package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/xunlun01?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败:", err.Error())
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	result01, err := tx.Exec("update account set money = ? where id = ?", 5000, 1001)
	result02, err := tx.Exec("update account set money = ? where id = ?", 1000, 1003)

	raf01, err := result01.RowsAffected()
	fmt.Println(err)
	raf02, err := result02.RowsAffected()
	fmt.Println(err)

	if raf01 == 1 && raf02 == 1 {
		tx.Commit()
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
		tx.Rollback()
	}
}
