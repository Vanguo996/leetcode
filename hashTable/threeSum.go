package main

func threeSum(nums []int) [][]int {
	// 输出不重复的解
	// [-1, -1, 2]  [-1, 2, -1]  [2, -1, -1]都是重复的
	// 需要去重和排序: 用map统计出现的次数，对key数组进行排序，
	quicksort(nums, 0, len(nums) - 1)
	var res [][]int
	//穷举threeSum的第一个数
	for i := 0; i < len(nums); i++ {
		//找到第一个数，穷举剩下的二元组
		tmp := twoTargetSum(nums, i + 1, 0 - nums[i])
		for _,v := range tmp {
			v = append(v, nums[i])
			res = append(res, v)
		}
		//跳过第一个数字的重复情况，
		for (i < len(nums) - 1) && (nums[i] == nums[i+1]) {
			i++
		}
	}


	return res

}

func twoTargetSum(nums []int, start, target int )  [][]int{
	//左指针从start开始, 找到不重复的二元组
	l, r := start, len(nums) - 1
	var res [][]int
	for l < r {
		sum := nums[l] + nums[r]
		lo , hi := nums[l], nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			res = append(res, []int{ nums[l], nums[r] })
			//判断重复
			for (l < r) && (lo == nums[l]) {
				l++
			}
			for (l < r) && (hi == nums[r]) {
				r--
			}

		}
	}
	return res
}



func quicksort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(arr, l ,r)
	quicksort(arr, l, i - 1)
	quicksort(arr, i + 1, r)
}

func partition(arr []int, l, r int) int{
	cur := l
	for i := l; i < r; i++{
		//当前值
		if arr[i] < arr[r]{
			arr[i], arr[cur] = arr[cur], arr[i]
			cur++
		}
	}
	arr[cur], arr[r] = arr[r], arr[cur]
	return cur
}