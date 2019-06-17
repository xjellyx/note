package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/srlemon/note/sql_/obj"
)

type student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	var db *obj.DB
	var (
		s = []interface{}{}
		//c=make(map[string]bool)
	)
	db = new(obj.DB)
	db.TableName = "users"
	db.Open()
	defer db.Close()
	b, err := db.Columns()
	fmt.Println(b, err)
	for _, v := range b {
		s = append(s, v["column_name"])
	}
	fmt.Println(s)

}
