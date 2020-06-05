package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {

	fmt.Println(reorganizeString("vvvlo"))

}
func reorganizeString(s string) string {
	var (
		arr []int
		lt  = list.New()
		res string
		rep = 0
	)
	if len(s) == 0 {
		return ""
	}
	for _, v := range s {
		arr = append(arr, int(v))
	}
	lt.PushBack(string(arr[0]))
	sort.Ints(arr)
	initLen := len(arr)
	for i := 1; i < len(arr); i++ {
		var (
			try bool
		)
		if lt.Back().Value != string(arr[i]) {
			lt.PushBack(string(arr[i]))
		} else if lt.Front().Value != string(arr[i]) {
			lt.PushFront(string(arr[i]))
		} else {
			var (
				data = lt.Back()
			)
			for data != nil {
				if i >= initLen {
					if data.Prev() == nil {
						fmt.Println(data, string(arr[i]))
					} else {
						fmt.Println(data, data.Prev(), data.Prev().Next(), string(arr[i]))
					}

				}
				if data.Prev() != nil && data.Prev().Value != string(arr[i]) && data.Prev().Prev() != nil && data.Prev().Prev().Value != string(arr[i]) {
					lt.InsertAfter(string(arr[i]), data.Prev().Prev())
					break
				} else {
					if !try {
						arr = append(arr, arr[i])
						try = true
						rep++
						if rep > initLen*2 {
							return ""
						}
					}
				}
				data = data.Prev()
			}
		}
	}
	var data = lt.Front()
	println(lt.Len())
	for data != nil {
		res += data.Value.(string)
		data = data.Next()
	}
	if len(res) != len(s) {
		fmt.Println(res)
		return ""
	}
	return res
}
