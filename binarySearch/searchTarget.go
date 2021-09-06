package binarySearch

/* 33搜索旋转排序数组

寻找有序数组中的目标值，如果没有目标值，返回即将要插入的索引值

*/

func search(nums []int, target int) int {
	left, right := 0,  len(nums) - 1
	// ans := len(nums)

	for left <= right {
		mid := left + (right - left) / 2
		if target > mid {
			// ans = mid
			left = mid + 1
		}else if target < mid {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}




/* 81搜索旋转排序数组

寻找有序数组中的目标值，如果没有目标值，返回即将要插入的索引值

*/

// func search(nums []int, target int) bool {

// }