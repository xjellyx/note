package main

import (
	"fmt"
	"sort"
)

// fibonacci_number 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var (
		fn = make([]int, n+1)
	)
	fn[0] = 0
	fn[1] = 1
	for i := 2; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	return fn[n]
}

// getMaximumGenerated
func getMaximumGenerated(n int) int {
	var (
		fn = make([]int, n+1)
	)
	if n == 0 {
		return 0
	}
	fn[0] = 0
	fn[1] = 1
	if n == 1 {
		return 1
	}
	for i := 0; i <= n; i++ {
		o := 2 * i
		j := o + 1
		if o >= 2 && o <= n {
			fn[o] = fn[i]
		}
		if j >= 2 && j <= n {
			fn[j] = fn[i] + fn[i+1]
		}
	}

	sort.Ints(fn)
	return fn[n]
}

// maxSubArray
func maxSubArray(nums []int) int {
	var (
		l   = len(nums)
		dp  = make([]int, l)
		max = nums[0]
	)
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		a := dp[i-1] + nums[i]
		b := nums[i]
		if a > b {
			dp[i] = a
		} else {
			dp[i] = b
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

// tribonacci 第n个泰波那契数
func tribonacci(n int) int {
	var (
		fn = make([]int, n+1)
	)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fn[0] = 0
	fn[1] = 1
	fn[2] = 1
	for i := 3; i <= n; i++ {
		fn[i] = fn[i-3] + fn[i-2] + fn[i-1]
	}
	return fn[n]
}

// climbStairs 爬楼梯
func climbStairs(n int) (res int) {
	var (
		dp = make(map[int]int)
	)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// minCostClimbingStairs 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		if cost[0] > cost[1] {
			return cost[1]
		} else {
			return cost[0]
		}
	}
	var (
		all []int
		dp  = make([]int, len(cost)+2)
	)
	all = append(all, 0) // 第0步
	all = append(all, cost...)
	all = append(all, 0) // 越过最后一步

	for i := 2; i < len(all); i++ {
		a := dp[i-2] + all[i-2]
		b := dp[i-1] + all[i-1]
		if a < b {
			dp[i] = a
		} else {
			dp[i] = b
		}
	}
	return dp[len(all)-1]
}

// rob house robber 打家劫舍
func rob(nums []int) int {
	var (
		l  = len(nums)
		dp = make([]int, l)
	)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	dp[0] = nums[0]

	if l >= 2 {
		if nums[1] < nums[0] {
			dp[1] = nums[0]
		} else {
			dp[1] = nums[1]
		}
	}

	for i := 2; i < l; i++ {
		a := dp[i-2] + nums[i]
		b := dp[i-1]
		if a > b {
			dp[i] = a
		} else {
			dp[i] = b
		}
	}
	return dp[l-1]
}

// deleteAndEarn 删除并获得点数
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

func max(val int, val2 int) int {
	if val > val2 {
		return val
	}
	return val2
}

// rob2 打家劫舍二
func rob2(nums []int) int {
	var (
		l     = len(nums)
		dp    = make([]int, l-1)
		dp2   = make([]int, l-1)
		nums2 = nums[1:]
		nums3 = nums[:l-1]
	)
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp[0] = nums2[0]
	if len(nums2) >= 2 {
		dp[1] = max(nums2[0], nums2[1])

	}
	for i := 2; i < l-1; i++ {
		dp[i] = max(dp[i-2]+nums2[i], dp[i-1])
	}
	dp2[0] = nums3[0]
	if len(nums3) >= 2 {
		dp2[1] = max(nums3[0], nums3[1])
	}
	for i := 2; i < l-1; i++ {
		dp2[i] = max(dp2[i-2]+nums3[i], dp2[i-1])
	}
	return max(dp[l-2], dp2[l-2])
}

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

func jump(nums []int) int {
	var (
		l           = len(nums)
		steps       int
		end         int
		maxPosition int
	)

	for i := 0; i < l; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
func main() {
	fmt.Println(rob2([]int{1, 2, 3, 1}))
}
