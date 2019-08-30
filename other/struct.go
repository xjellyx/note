package main

import "fmt"

type Student struct{
	Name string // 姓名
	Age int // 年龄
	Id string // 学号
	Class string // 班级
}

func main(){
	var s = new(Student)
	s.Name = "詹姆斯"
	s.Age = 35
	s.Id = "23"
	s.Class = "湖人"

	fmt.Println(s)
}
