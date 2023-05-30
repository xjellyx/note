package main

import "fmt"

// ShellSort 希尔排序
func ShellSort(arr []int) {
	length := len(arr)
	// 初始步长为数组长度的一半
	gap := length / 2
	for gap > 0 {
		// 插入排序
		for i := gap; i < length; i++ {
			j := i
			for j >= gap && arr[j-gap] > arr[j] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
		gap /= 2
	}
}

func main() {
	arr := []int{7, 1, 3, 6, 5, 4, 2, 9, 8}
	fmt.Println("排序前:", arr)
	ShellSort(arr)
	fmt.Println("排序后:", arr)
}
