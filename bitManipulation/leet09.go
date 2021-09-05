package bitManipulation

func isPalindrome(x int) bool {
	//整数反转 + 判断
	if x < 0 {
		return false
	}
	var y int64
	flag := int64(x)
	for x != 0 {
		y = y * 10 + int64(x % 10)
		x = x / 10
	}
	return flag == y
}
