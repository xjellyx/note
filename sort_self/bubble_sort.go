package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 记录是否交换
		flag := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		// 如果没有交换，说明以是有序
		if !flag {
			break
		}
	}
}

func main() {
	arr := []int{7, 1, 3, 6, 5, 4, 2, 9, 8}
	fmt.Println("排序前:", arr)
	BubbleSort(arr)
	fmt.Println("排序后:", arr)
}
