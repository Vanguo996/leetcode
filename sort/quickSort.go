package main

import (
	"fmt"
	"time"
)

func main() {

	r := 10
	for i := 0; i < r; i++ {
		go printGo(i)
	}
}

func printGo(i int) {

	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println("hello %d", i)
}

// func printGo(i int) {
// sleep(time.Second * i)

// time.Sleep()

// }

// 输入数组
//func quicksort(arr []int, l int, r int) {
//	if l < r{
//		p := partition(arr, l, r) //返回等于区域
//		quicksort(arr, l, p[0] - 1)
//		quicksort(arr, p[1] + 1, r)
//	}
//}
//
//func partition(arr []int, l int, r int) []int{
//	less := l - 1
//	more := r
//	cur := l
//	for cur < more {
//		if arr[cur] < arr[r] {
//			//小于基准值，区域右移，和cur交换，cur右移
//			less++
//			swap(arr, less, cur)
//			cur++
//		} else if arr[cur] > arr[r]{
//			more--
//			swap(arr, cur, more)
//		} else {
//			cur++
//		}
//	}
//	swap(arr, more, r) //基准值交换
//	p := []int{less + 1, more}
//	return p
//}
//
//func swap(arr []int, l int, r int) {
//	arr[l], arr[r] = arr[r], arr[l]
//}
//

func quicksort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(arr, l, r)
	quicksort(arr, l, i-1)
	quicksort(arr, i+1, r)
}

func partition(arr []int, l, r int) int {
	cur := l
	for i := l; i < r; i++ {
		//当前值
		if arr[i] < arr[r] {
			arr[i], arr[cur] = arr[cur], arr[i]
			cur++
		}
	}
	arr[cur], arr[r] = arr[r], arr[cur]
	return cur
}

//2021.2.20

func quickSort2(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, l int, r int) {
	//递归条件注意
	if l >= r {
		return
	}
	i := partition(arr, l, r)
	sort(arr, l, i-1)
	sort(arr, i+1, r)
}

func partition2(arr []int, l int, r int) int {
	cur := l
	pivot := arr[r]
	for i := l; i < r; i++ {
		if arr[i] < pivot {
			arr[i], arr[cur] = arr[cur], arr[i]
			cur++
		}
	}
	arr[cur], arr[r] = arr[r], arr[cur]
	return cur
}

// func main() {
// 	var arr = []int{23, 4, 3, 6, 1, 56, 21, 33, 9}
// 	//sort.Ints(arr)
// 	quickSort2(arr)
// 	fmt.Print(arr)

// }
