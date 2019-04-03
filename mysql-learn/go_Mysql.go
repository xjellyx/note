package main

import (
	"fmt"
	obj "github.com/LnFen/note/mysql-learn/obj"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	var db *obj.Mysql_db
	var (
		s = []interface{}{}
		//c=make(map[string]bool)
	)
	db = new(obj.Mysql_db)
	db.TableName = "users"
	db.MysqlOpen()
	defer db.MysqlClose()
	b, err := db.Clounms()
	fmt.Println(b, err)
	for _, v := range b {
		s = append(s, v["column_name"])
	}
	fmt.Println(s)

}
