package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	name = "sdsd"
	t := reflect.TypeOf(name)
	fmt.Println(t)
}
