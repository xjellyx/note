package main

import (
	"fmt"
	"sort"
)

// sortedSquares 有序数组的平方
func sortedSquares(nums []int) []int {
	var (
		newNums = make([]int, len(nums))
	)
	for i, v := range nums {
		newNums[i] = v * v
	}

	sort.Ints(newNums)
	return newNums
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	l := len(nums)
	left, right := 0, 0
	for right < l {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	var res []int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if numbers[i]+numbers[j] == target {
				res = append(res, i+1, j+1)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}
