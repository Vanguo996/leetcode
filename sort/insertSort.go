package main

// import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] > arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func insertSort2(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// func main() {
// 	arr := []int {7,5,3,9,2,1,18,15,13}
// 	res := insertSort2(arr)
// 	fmt.Print(res)
// }
