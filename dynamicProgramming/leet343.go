package dynamicProgramming

func integerBreak(n int) int {
	//这一一道DP题目，将数字拆分多个数字之和，求分解出来的数字最大是多少
	//这一题动态转移方程是dp[i] = max(dp[i], j * (i - j), )
	// 一个数分解成 i , i - j,或者分解成j 和 更多分解数 -> dp[i - j]

	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j * max(dp[i - j], i - j))
		}
	}
	return dp[n]

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
