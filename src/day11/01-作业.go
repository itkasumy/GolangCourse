package main

import "fmt"

type Database interface {
	insert()
	update()
	query()
	delete()
}

type Mysql struct {
	name string
}

func (msql Mysql) insert() {
	fmt.Println(msql.name, "插入数据成功")
}

func (msql Mysql) update() {
	fmt.Println(msql.name, "修改数据成功")
}

func (msql Mysql) query() {
	fmt.Println(msql.name, "查询数据成功")
}

func (msql Mysql) delete() {
	fmt.Println(msql.name, "删除数据成功")
}

type Oracle struct {
	name string
}

func (orc Oracle) insert() {
	fmt.Println(orc.name, "插入数据成功")
}

func (orc Oracle) update() {
	fmt.Println(orc.name, "修改数据成功")
}

func (orc Oracle) query() {
	fmt.Println(orc.name, "查询数据成功")
}

func (orc Oracle) delete() {
	fmt.Println(orc.name, "删除数据成功")
}

type Project struct {
	name string
	db Database
}

func (pro Project) testDataBase() {
	fmt.Println(pro.name, "开始工作了...")
	pro.db.insert()
	pro.db.delete()
	pro.db.update()
	pro.db.query()
}

func main() {
	//定义一个DataBase接口，4个方法：insert(),update(),query(),delete()
	//定义两个结构体实现该接口：Mysql和Oracle
	//定义一个Project结构体：两个字段：name表示项目名字，DataBase表示项目用的数据库，定义一个方法：testDataBase(),测试DataBase的4个方法。
	pro01 := Project{"xunlun", Mysql{"mysql"}}
	pro01.testDataBase()

	pro02 := Project{"xuexi", Oracle{"oracle"}}
	pro02.testDataBase()
}
