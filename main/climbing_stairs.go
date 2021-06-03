package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
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
