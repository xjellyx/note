package main

import "fmt"

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 记录最小元素的index
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 把最小元素放入当前位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{7, 1, 3, 6, 5, 4, 2, 9, 8}
	fmt.Println("排序前:", arr)
	SelectionSort(arr)
	fmt.Println("排序后:", arr)
}
