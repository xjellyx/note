package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"strconv"
)

// User
type User struct {
	Id       int64  `json:"id"`
	Name     string `xorm:"varchar(25) notnull unique "`
	Age      int    `json:"age" xorm:"int(3)"`
	Username string `json:"username" xorm:"varchar(25) notnull unique"`
	ClassId  string `json:"class_id" xorm:"index"`
}

var (
	engine *xorm.Engine
)

func main() {
	var (
		err error
	)
	// 连接数据库
	if engine, err = xorm.NewEngine("mysql", "business:business@127.0.0.1:3306/business?charset=utf8"); err != nil {
		panic(err)
	}

	// 创建user表
	if err = engine.CreateTables(&User{}); err != nil {
		panic(err)
	}
	// 根据struct中的tag来创建唯一索引
	_ = engine.CreateUniques(&User{})
	// 根据struct中的tag来创建索引
	_ = engine.CreateIndexes(&User{})
	// 打印sql语句
	engine.ShowSQL(true)
	// 同步表的数据
	//if err  = engine.Sync2(new(User));err!=nil{
	//	panic(err)
	//}

	// 插入一条数据
	for i := 0; i < 10; i++ {
		if err = InsertUser(&User{
			Name:     "张三" + strconv.Itoa(i),
			Age:      18,
			Username: "asfdasfas" + strconv.Itoa(i),
			ClassId:  "3217498375904",
		}); err != nil {
			log.Fatalln("[InsertUser] err: ", err)
		}
	}

	//// 获取信息
	//if _d,_err := GetUserData("zhangsan999");_err!=nil{
	//	log.Errorln("[GetUserData] err: ",_err)
	//}else {
	//	log.Infoln("[GetUserData] data: ",_d)
	//}
	//
	//// 更新数据
	//if err = UpdateUser("zhangsan999", struct {
	//	Name string
	//	Age  int
	//}{Name:"张三更新" , Age:19 });err!=nil{
	//	log.Errorln("[UpdateUser] err: ",err)
	//}
	//
	//// 获取信息
	//if _d,_err := GetUserData("zhangsan999");_err!=nil{
	//	log.Errorln("[GetUserData] err: ",_err)
	//}else {
	//	log.Infoln("[GetUserData] data: ",_d)
	//}
	//
	//// 删除数据
	//if err = DeleteUser("zhangsan999");err!=nil{
	//	log.Errorln("[DeleteUser] err: ", err )
	//}
	//
	//// 获取信息
	//if _d,_err := GetUserData("zhangsan999");_err!=nil{
	//	log.Errorln("[GetUserData] err: ",_err)
	//}else {
	//	log.Infoln("[GetUserData] data: ",_d)
	//}

}

// InsertUser 插入用户数据
func InsertUser(u *User) error {
	var (
		d   int64
		err error
	)

	if d, err = engine.Insert(u); err != nil {
		return err
	}
	fmt.Println("aaaaaaaaaaa", d)
	return nil
}

// GetUserData 获取用户信息
func GetUserData(username string) (ret *User, err error) {
	var (
		has  bool
		data = new(User)
	)
	if has, err = engine.Table(&User{}).Where("username=?", username).Get(data); err != nil {
		return nil, err
	} else if !has {
		err = xorm.ErrNotExist
		return
	}
	//
	ret = data
	return
}

// UpdateUser 更新用户信息
func UpdateUser(username string, updateForm struct {
	Name string
	Age  int
}) (err error) {
	if _, err = engine.Table(&User{}).Where("username = ?", username).Update(map[string]interface{}{
		"name": updateForm.Name,
		"age":  updateForm.Age,
	}); err != nil {
		return
	}
	return
}

// DeleteUser 删除用户数据
func DeleteUser(username string) (err error) {
	if _, err = engine.Where("username = ?", username).Delete(&User{}); err != nil {
		return err
	}
	return
}
