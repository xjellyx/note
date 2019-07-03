package main

/*
#include <stdlib.h>
// 接口形式1
void SayHello(char* s);
// 接口形式2
void SayHello2(_GoString_ s);

*/
import "C"

import (
	"fmt"
)

func main() {
	C.SayHello(C.CString("Hello, World\n"))
	C.SayHello2("hello,cgo !!! \n")
	fmt.Println(int(C.random()))
	fmt.Println(1 << 32)
	a := 3
	if a != 0 && a != 2 {
		println("aaaaaaa")
	} else {
		println(a)
	}
	var s *int32
	fmt.Println(int(*s))
}

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}

//export SayHello2
func SayHello2(s string) {
	fmt.Println(s)
}
