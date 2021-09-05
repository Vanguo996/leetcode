package main

import "fmt"

func main() {
	s := "abcabcdebb"
	fmt.Print(foo(s))
}


func foo(s string) int{
	maxLen := 0
	dict := map[byte]bool{}
	str := []byte(s)
	left := 0

	for right, val := range str {
		for dict[val] {
			_, ok := dict[str[left]]
			if ok {
				dict[str[left]] = false
			}
			left++
		}
		dict[val] = true
		maxLen = max(maxLen, right - left + 1)
	}
	return maxLen

}

func max(a, b int)  int{
	if a < b {
		return b
	}
	return a
}