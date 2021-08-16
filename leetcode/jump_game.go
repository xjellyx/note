package main

import "fmt"

func canJump(nums []int) bool {
	l := len(nums)
	reach := 0
	for i := 0; i < l; i++ {
		if i > reach {
			return false
		}
		if nums[i]+i > reach {
			reach = nums[i] + i
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{3, 1, 4, 2, 0, 9}))
}
