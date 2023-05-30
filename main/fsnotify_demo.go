package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"unsafe"
)

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func (p Pill) String() string {
	return strconv.Itoa(int(p))
}

type QQ struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var q = QQ{
		Name: "wwwwwwww",
		Age:  111111,
	}
	var aa = unsafe.Pointer(&q)
	d := (*QQ)(atomic.LoadPointer(&aa))
	d.Name = "我改了名称，下面也被改了"
	fmt.Println(*d)
	fmt.Println(q)
}
