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

func min(val int, val2 int) int {
	if val < val2 {
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

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	maxVal := nums[0]
	if l == 1 {
		return maxVal
	}
	for i := 0; i < l; i++ {
		maxVal = max(maxVal, nums[i])
		v := nums[i]

		for j := i + 1; j <= l; j++ {
			if j < l {
				v += nums[j]
			} else {
				for m := 0; m < i; m++ {
					v += nums[m]
					maxVal = max(maxVal, v)
				}
			}

			maxVal = max(maxVal, v)
		}

	}
	return maxVal
}

// maxProfit 买卖股票最大收益
func maxProfit(nums []int) int {
	var (
		minV = nums[0]
		maxV = 0
		l    = len(nums)
	)
	for i := 0; i < l; i++ {
		maxV = max(maxV, nums[i]-minV)
		minV = min(minV, nums[i])
	}
	return maxV
}

// maxProfit2 买卖股票最大收益 2
func maxProfit2(nums []int) int {
	var (
		sum int
		l   = len(nums)
	)
	for i := 1; i < l; i++ {
		val := nums[i] - nums[i-1]
		if val > 0 {
			sum += val
		}
	}
	return sum
}

// generateParenthesis 括号生成
func generateParenthesis(n int) []string {

	res := make([]string, 0, 16)
	var dfs func(path string, left, right int)
	dfs = func(path string, left, right int) {
		if left > n || right > left {
			return
		}
		if len(path) == n*2 {
			res = append(res, path)
			return
		}
		dfs(path+"(", left+1, right)
		dfs(path+")", left, right+1)
	}
	dfs("", 0, 0)
	return res
}

//  杨辉三角
func generate(numRows int) [][]int {
	var (
		dp = make([][]int, numRows)
	)
	if numRows == 1 {
		return [][]int{{1}}
	}
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	dp[1][0] = 1
	dp[1][1] = 1

	for i := 2; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}

// 杨辉三角 II
func getRow(rowIndex int) []int {
	var (
		dp = make([][]int, rowIndex+1)
	)
	if rowIndex == 0 {
		return []int{1}
	}
	for i := 0; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	dp[1][0] = 1
	dp[1][1] = 1

	for i := 2; i < rowIndex+1; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp[rowIndex]
}

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

func getMaxLen(nums []int) int {
	l := len(nums)
	var (
		maxVal = 0
		maxL   = 0
		m      int
	)
	if l == 1 {
		return 1
	}

	for i := 0; i < l; i++ {
		v := nums[i]
		m = 1
		maxVal = max(maxVal, v)
		for j := i + 1; j < l; j++ {
			v *= nums[j]
			if maxVal < v {
				m++
			}
			maxVal = max(maxVal, v)
		}
		maxL = max(maxL, m)
	}

	return maxL
}

// 1014. 最佳观光组合
func maxScoreSightseeingPair(values []int) int {
	var (
		l    = len(values)
		maxV = values[0] + 0
		res  = 0
	)
	for i := 1; i < l; i++ {
		fmt.Println(res, maxV, values[i]-i, i)
		res = max(res, maxV+values[i]-i)
		maxV = max(maxV, values[i]+i)
	}
	return res
}

func maxProfit3(prices []int, fee int) int {
	var (
		dpIn  = make([]int, len(prices))
		dpOut = make([]int, len(prices))
	)
	dpIn[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		dpIn[i] = max(dpIn[i-1], dpOut[i-1]-prices[i])
		dpOut[i] = max(dpOut[i-1], dpIn[i-1]+prices[i]-fee)

	}
	return dpOut[len(prices)-1]
}

func test(s string) (res string) {
	var (
		sb   = []byte(s)
		l    = len(sb)
		slow = 0
	)
	for fast := 0; fast < l; fast++ {
		if sb[fast] != '#' {
			sb[slow] = sb[fast]
			slow++
		} else if slow > 0 {
			slow--
		}
	}

	res = string(sb[:slow])
	return res
}

func sortedSquares2(nums []int) []int {
	var (
		l       = len(nums)
		slow    = l - 1
		fast    = 0
		index   = l - 1
		newNums = make([]int, l)
	)
	for fast <= slow {
		v1 := nums[fast] * nums[fast]
		v2 := nums[slow] * nums[slow]
		if v1 > v2 {
			newNums[index] = v1
			fast++
		} else {
			newNums[index] = v2
			slow--
		}
		index--
	}

	return newNums
}

func main() {
	fmt.Println(sortedSquares2([]int{-5, -4, -2, -1, 0, 10}))

}
