package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p, e := bcrypt.GenerateFromPassword([]byte("123456789"), bcrypt.DefaultCost)
	fmt.Println(string(p), e)
	e = bcrypt.CompareHashAndPassword(p, []byte("123456789"))
	fmt.Println(e)
}
