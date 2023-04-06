package main

import "fmt"

func main() {
	var (
		nums = []int{3, 1, 5, 3, 6, 1, 4, 6, 78, 45, 13, 64}
	)
	//bubbleSort(nums)
	//insertSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}

func quickSort(nums []int, left, right int) {
	for left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		left = mid + 1
	}
}

func partition(nums []int, left, right int) int {
	// choose a pivot

	// 改变的地方
	pivot := threeSumMedian(nums[left], nums[(left+right)/2], nums[right])
	//pivot := nums[left]
	nums[left], pivot = pivot, nums[left]
	// 改变结束

	for left < right {
		for pivot <= nums[right] && left < right {
			right -= 1
		}
		nums[left] = nums[right]

		for nums[left] <= pivot && left < right {
			left += 1
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left

}

// input 10 20 30 ---> return 20 ; input 10 10 11 --> return 10
func threeSumMedian(a, b, c int) int {

	if a < b {
		a, b = b, a
	}
	// a  > b

	if c > a {
		return a
	} else {
		if c > b {
			return c
		} else {
			return b
		}
	}
}
