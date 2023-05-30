package main

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// 待插入的数据
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			// 元素往后移动
			arr[j+1] = arr[j]
		}
		// 插入元素
		arr[j+1] = temp
	}
}

func main() {
	arr := []int{7, 1, 3, 6, 5, 4, 2, 9, 8}
	fmt.Println("排序前:", arr)
	InsertionSort(arr)
	fmt.Println("排序后:", arr)
}
