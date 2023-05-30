package main

func numWays(n int) (ret int) {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var (
		dp = make([]int, n+1)
	)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
