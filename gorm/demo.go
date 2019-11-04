package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	db, err := gorm.Open("mysql", `business:business@tcp(127.0.0.1:33306)/business`)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	now := time.Now()
	time.Sleep(time.Second * 6)
	now1 := time.Now()
	fmt.Println(now1.After(now))

}
