package main

import (
	"fmt"
	"github.com/srlemon/note/variable"
)

type aa struct {
	A string
}

func (a *aa) Test() {
	fmt.Println(a.A)
}

func main() {
	var (
		a *aa
	)
	defer variable.PanicRecoverError(variable.Log)
	a.Test()

}
