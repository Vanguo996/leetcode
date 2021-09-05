package main

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
//但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
func findRepeatNumber(nums []int) int {
	//注意，因为是0- n -1范围之内，所以创建一个布尔类型数组
	res := [100000]bool{}
	for _, v := range nums{
		if res[v] {
			return v
		}
		res[v] = true
	}

	var sum int

	sum = 1
	print(sum)
	return 0

}
