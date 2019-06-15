package main

import (
	"fmt"
	"reflect"
)

func main() {
	tonydon := &user1{"TangXiaodong", 100, "0000123"}
	object := reflect.ValueOf(tonydon)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	}
	tonydon.SayHello()
	v := object.MethodByName("SayHello")
	v.Call([]reflect.Value{})
}

type user1 struct {
	Name string
	Age  int
	Id   string
}

func (u *user1) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}
