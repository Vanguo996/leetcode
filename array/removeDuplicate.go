package main

func remove(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var i = 0
	last := nums[len(nums) - 1]
	for i = 0; i < len(nums); i++ {
		//记录最后一个值，
		if nums[i] == last {
			break
		}
		if nums[i] == nums[i + 1] {
			//两个数相等移动后面多余的
			removeElement(nums, i + 1, nums[i])
		}
	}
	return i + 1

}

func removeElement(nums []int, start, val int) int{
	//移动零
	//cur指向val，i负责寻找与目标值不相等的数
	j := start
	for i := start; i < len(nums); i++ {
		//nums[i] == val 只移动i
		if nums[i] != val {
			//如果i == j 不交换，移动j
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

// golang语言原地去重

func removeDup(nums []int) []int {
	i := 0
	for j := 1; j < len(nums); j++{
		if nums[i] != nums[j]{
			i++
			// 赋值或者交换
			nums[i] = nums[j]

		}
	}
	return nums[:i+1]

}











func main() {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	remove(arr)

}



