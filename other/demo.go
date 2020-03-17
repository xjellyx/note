package main

import (
	"fmt"
	"time"
)

type MapSorter []Item

type Item struct {
	Key string
	Val int64
}

func NewMapSorter(m map[string]int64) MapSorter {
	ms := make(MapSorter, 0, len(m))

	for k, v := range m {
		ms = append(ms, Item{k, v})
	}

	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Val < ms[j].Val // 按值排序
	//return ms[i].Key < ms[j].Key // 按键排序
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	go sayHello()
	time.Sleep(time.Second)
}

func sayHello() {
	fmt.Println("Hello Goroutine !!!")
}
