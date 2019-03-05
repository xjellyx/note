package main

import (
	"fmt"
)

type D struct {
	Age int
	Name string
}
func main()  {
    var a =new(D)
    a.Name="ssssssss"
    a.Age=16
    fmt.Println(a)

    a.Name="ddddddddd"
    fmt.Println(a)

}

