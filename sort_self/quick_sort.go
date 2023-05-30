package main

import "fmt"

// QuickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left < right {
		// 分治
		pos := partition(arr, left, right)
		// 左半部分递归排序
		QuickSort(arr, left, pos-1)
		// 右半部分递归排序
		QuickSort(arr, pos+1, right)
	}
}

// 分治
func partition(arr []int, left, right int) int {
	// 取首元素作为基准元素
	pivot := arr[left]
	for left < right {
		// 找到第一个大于基准元素的位置
		for left < right && arr[right] >= pivot {
			right--
		}
		// 右侧元素放入左侧
		arr[left] = arr[right]
		// 找到第一个小于基准元素的位置
		for left < right && arr[left] <= pivot {
			left++
		}
		// 左侧元素放入右侧
		arr[right] = arr[left]
	}
	// 将基准元素放入中间位置
	arr[left] = pivot
	// 返回中间位置
	return left
}

func main() {
	arr := []int{7, 1, 3, 6, 5, 4, 2, 9, 8}
	fmt.Println("排序前:", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("排序后:", arr)
}
