package bitManipulation

//整数反转
func reverse(x int) int {
	y := 0
	for x != 0 {
		// 如果越界返回0
		if (y > 214748364 || y < -214748364) {
			return 0
		}
		y = y * 10 + x % 10
		x = x / 10
	}
	return y
}