package main

import "fmt"

func main() {
	arr := []int{5, 8, 1, 3, 9, 6}
	fmt.Println("Before sorting:", arr)
	arr = mergeSort(arr)
	fmt.Println("After sorting:", arr)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}

	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]

	leftArr := mergeSort(left)
	rightArr := mergeSort(right)

	mergedArr := merge(leftArr, rightArr)

	return mergedArr
}

func merge(left []int, right []int) []int {
	var i, j int
	leftLen := len(left)
	rightLen := len(right)
	mergedArr := make([]int, leftLen+rightLen)

	for i < leftLen && j < rightLen {
		if left[i] < right[j] {
			mergedArr[i+j] = left[i]
			i++
		} else {
			mergedArr[i+j] = right[j]
			j++
		}
	}

	for i < leftLen {
		mergedArr[i+j] = left[i]
		i++
	}

	for j < rightLen {
		mergedArr[i+j] = right[j]
		j++
	}

	return mergedArr
}
