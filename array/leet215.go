package main

import (

)

func findKthLargest(nums []int, k int) int {
	//sort.Ints(nums)
	//return nums[len(nums) - k]

	//快排的思想
//	left := 0
//	right := len(nums) - 1
//
//	for {
//		if left >= right {
//			return nums[right]
//		}
//		p := partition(nums, left, right)
//		if p + 1 == k {
//			return nums[p]
//		} else if p + 1 < k {
//			left = p + 1
//		} else {
//			right = p - 1
//		}
//	}
//}
//
//func partition(nums []int, left, right int) int {
//	for i := left; i < right; i++ {
//		if nums[i] < nums[right] {
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//		}
//	}
//	nums[left], nums[right] = nums[right], nums[left]
//	return left

//堆排序

	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	//[0, 5, 3, 2, 1, 4]
	for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]

}

func buildMaxHeap(a []int, size int)  {
	for i := size / 2; i >= 0; i-- {
		maxHeapify(a, i, size)
	}
}

func maxHeapify(a []int, i, size int) {
	l, r, largest := i * 2 + 1, i * 2 + 2, i
	if l < size && a[l] > a[largest] {
		largest = l
	}
	if r < size && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, size)
	}
}



//rand.Seed(time.Now().UnixNano())
//return quickSelect(nums, 0, len(nums) - 1, len(nums) - k)

//func quickSelect(a []int, l, r, index int) int {
//	q := randomPartition(a, l, r)
//	if q == index {
//		return a[q]
//	} else if q < index {
//		return quickSelect(a, q + 1, r, index)
//	}
//	return quickSelect(a, l, q - 1, index)
//}
//
//func randomPartition(a []int, l, r int) int {
//	i := rand.Int() % (r - l + 1) + 1
//	a[i], a[r] = a[r], a[i]
//	for i := l; i < r; i++{
//		if a[i] <= a[r] {
//			a[i], a[l] = a[l], a[i]
//			l++
//		}
//	}
//	a[l], a[r] = a[r], a[l]
//	return l
//}











