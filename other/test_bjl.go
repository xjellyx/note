package main

import (
	"sync"
)

var (
	a, b = new(sync.Map), new(sync.Map)
)

type ddd struct {
	M int
	N int
}

func init() {
	a.Store(1, "10")
	var (
		demo = new(ddd)
	)
	demo.M = 22
	demo.N = 99
	b.Store("10", demo)
}

func main() {

}

func sddd(c **ddd) {
	(*c).N -= 1
}
