package main

import (
	"fmt"
)

func main() {
	var (
		arr = []int{3, 1, 45, 321, 13, 34543, 234}
	)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		fmt.Println("aaaaaaaa", mid)
		quickSort(arr, left, mid-1)
		quickSort(arr, mid+1, right)
	}

}

func partition(arr []int, left, right int) int {
	var (
		key = arr[left]
	)
	for left < right {
		// 找到右边比基准值小的数据
		for left < right && arr[right] > key {
			right--
		}
		arr[left] = arr[right]
		// 找到左边比基准值大的数据
		for left < right && arr[left] < key {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	//fmt.Println(key, left, right)
	return left
}

func FastSort(arr []int, left, right int) {
	if left < right {
		//找到中心轴排好序的位置
		pivot := partition1(arr, left, right)
		//对低子表递归排序
		FastSort(arr, left, pivot-1)
		//对高子表递归排序
		FastSort(arr, pivot+1, right)
	}
}

// 对数据元素进行左右调整
// 小于中心值则向左边调换
// 大于中心值则向右边调换
// 借用left, right指针,一个左,一个右.
func partition1(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] > pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] < pivot {
			left++
		}
		arr[right] = arr[left]
	}
	//跳出for表示left,right指针重合.
	arr[left] = pivot
	fmt.Printf("pivot:%d,left:%d,right:%d,%v\n", pivot, left, right, arr)
	return left
}
