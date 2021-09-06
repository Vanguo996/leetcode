package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int{
	m := make(map[int]int)
	for i := 0; i < len(m); i++ {
		another := target - nums[i]
		// 查找是否key是否存在
		// 如果存在，取出下标
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		// 如果不存在放入下标值
		m[nums[i]] = i
	}
	return nil
}


// find 统计数字出现的次数
func find(nums []int) map[int]int {
	m := make(map[int]int)

	for _, val := range nums{
		m[val]++
	}
	return m
}


type TwoSum struct {
	m map[int]int
}

func NewTwoSum() *TwoSum {
	return &TwoSum{
		m: make(map[int]int,0),
	}
}

func (s *TwoSum) add(number int) {
	// 记录number出现的次数
	s.m[number]++
}

//func (s *TwoSum) find(val int)  bool{
//
//}


//如果数组有序

func twoSum3(nums []int, target int) []int{
	sort.Ints(nums)
	left, right := 0, len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left, right}
		}
	}

	return nil
}

//重复元素，排序，双指针
func twoSumMultiple(nums []int, target int){
	l, r := 0, len(nums) - 1
	var res [][]int
	for l < r {
		sum := nums[l] + nums[r]
		//记录最初索引值
		lo, hi := nums[l], nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			res = append(res, []int{nums[l], nums[r]})
			//跳过重复元素
			for (l < r) && (nums[l] == lo) { l++ }
			for (l < r) && (nums[r] == hi) { r-- }
		}
	}
	fmt.Print(res)
}






func main() {
	arr := []int{3,0,-2,-1,1,2}
	res := threeSum(arr)
	fmt.Print(res)


}

// 2021.9.5 
func twoSumNew(arr []int, target int) []int{
	m := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		temp := target - arr[i]
		if _, ok := m[temp]; ok {
			return []int{m[temp], i}
		}
		m[arr[i]] = i
	}
	return nil	
}