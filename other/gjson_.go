package main

/*
#include <stdio.h>
int test(int a, int b){
return a + b;
}


*/
import "C"
import (
	"fmt"
	"github.com/tidwall/gjson"
)

var jstr = `{
  "name": {"first": "Tom", "last": {"Anderson": "sdhoaishd"}},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}`

func main() {
	r := gjson.Get(jstr, "name.last")
	fmt.Println(r.String())
	a := C.int(10)
	b := C.int(20)
	fmt.Println(C.test(a, b))

}
