package main

//在gin中连接和操作mysql与在其他框架中连接操作mysql没有什么区别，分为主要的几个步骤：

//1、引入mysql驱动程序：使用import将mysql驱动默认引入，具体语法如下：
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	//2、拼接链接字符：在程序中链接mysql，需要按照一定的规则进行用户名，密码等信息的组合。
	connStr := "root:123456@tcp(127.0.0.1:3306)/ginsql"
	//3、使用sql.Open创建数据库连接
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//4、操作数据库

	//a、 创建数据库表
	/*	_, err = db.Exec("create table person(" +
			"id int auto_increment primary key," +
			"name varchar(12) not null," +
			"age int default 1" +
			");")
		if err!= nil {
			log.Fatal(err.Error())
			return
		}*/

	//b、向数据库中插入数据
	_, err = db.Exec("insert into person(name ,age) values(?,?); ", "zhu", 24)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("插入数据成功")
	}

	//c、查询数据库记录
	rows, err := db.Query("select * from person ;")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
scan:
	if rows.Next() {
		//columns ,_ := rows.Columns()
		//fmt.Println(columns)
		person := new(Person)
		/*
			查询出来的row，只能用Scan方法获取值。
			而Scan方法必须按照查询出来的列，一一对应把参数给他塞进去，才能取到值。
			多一个少一个都会返回error
		*/
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}

}

type Person struct {
	Id   string
	Name string
	Age  string
}
