package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		list []int
	)

	for i := 1; i <= 100000000; i++ {
		list = append(list, i)
	}
	start := time.Now()
	fmt.Println(getIndex(list, 0, len(list)-1, 96454536))
	end := time.Now()
	fmt.Println(float64(end.UnixNano()-start.UnixNano()) / float64(time.Second))

}

func getIndex(list []int, left, right int, target int) (index int) {
	if left > right { // 找不到存在的数据, 返回-1
		return -1
	}
	mid := (left + right) / 2 // 每次的中间id
	if target == list[mid] {
		return mid
	} else if target > list[mid] {
		return getIndex(list, mid+1, right, target)
	} else {
		return getIndex(list, left, mid-1, target)
	}

}
