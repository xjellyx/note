package main

import (
	"fmt"
	"reflect"
)

type people struct {
	name string
	age  int
}

func test(data interface{}) {
	t := reflect.TypeOf(data)
	fmt.Println(t)

	v := reflect.ValueOf(data)
	fmt.Println(v)

	fmt.Println(v.Kind())

	fmt.Println(v.Interface())
	fmt.Println(v.Interface().(*people))
}
func testStruct(v interface{}) {
	val := reflect.ValueOf(v)
	kd := val.Kind()
	fmt.Println(val, kd)

	fields := val.NumField()
	fmt.Println(fields)
	for i := 0; i < fields; i++ {
		fmt.Println(val.Field(i).Kind())
	}

	fmt.Println("ssssssssssss")
	methods := val.NumMethod()

	fmt.Println(methods)
}
func main() {
	p := &people{
		name: "hklhaff",
		age:  18,
	}
	testStruct(*p)
}
