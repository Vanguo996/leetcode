package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


//3
//3 1 4 3
//4 2 5 3 1
//4 13 9 5 5

// 输入输出
func main() {
	//input := bufio.NewScanner(os.Stdin)
	//var s [][]string
	//
	//input.Scan()
	//t := input.Text()
	//i, _ := strconv.Atoi(t)
	//
	//for j := i ; j > 0; j--{
	//	input.Scan()
	//	str1 := strings.Split(input.Text(), " ")
	//	s = append(s, str1[1:])
	//}
	////foo(s)

	io()
}

func io() {
	input := bufio.NewScanner(os.Stdin)
	var s []string
	input.Scan()
	t := input.Text()
	i, _ := strconv.Atoi(t)

	for j := i ; j > 0; j-- {
		input.Scan()
		str1 := input.Text()
		s = append(s, str1)
	}

	foo(s)

}

func foo(s []string) {

	m := make(map[string]string)
	set := make(map[string]struct{})
	var tail string
	length := len(s)
	tmp := strings.Split(s[0], " ")
	head := tmp[0]
	tail = tmp[1]

	 //存入map中
	for _, v := range s {
		i := strings.Split(v, " ")
		m[i[0]] = i[1]
	}

	var ok bool

	//从尾部开始查找，传递给tail，如果head = tail那么结束查找。

	//
	set[head] = struct{}{}
	set[tail] = struct{}{}
	count := 1
	var res int
	for  {

		if tail, ok = m[tail]; ok {
			//查找到放入set中
			set[tail] = struct{}{}
			//每查询一次，count计数
			count++
		}

		if head == tail {
			res++
			//重新定位新的  head  and tail
			// 通过map和set比较来判断，从哪个新的开始
			//遍历set中的元素，如果不在map中，那么更新head 和  tail
			for v := range set {
				if key, exist := m[v]; !exist {
					head = key
					tail = m[key]
				}
			}
		}

		if count == length{
			fmt.Print(res)
			break
		}

	}

}





//func foo(s [][]string) bool{
//
//	for _, v := range s {
//		length := len(v)
//		for i := 0; i < length - 2; i++ {
//			//如果不是递减的，断点是i + 1
//			if v[i] < v[i + 1] {
//				if v[0] > v[length - 1] {
//					return false
//				} else {
//					return true
//				}
//			}
//		}
//		//无切分
//		return true
//	}
//
//}
//













//func asciiSum(s string) int {
//	var ret = 0
//	for _, b := range s {
//		ret += int(b)
//	}
//	return ret
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//func miniSum(s1 string, s2 string) int {
//	var (
//		len1, len2 = len(s1), len(s2)
//		dp         = make([][]int, len1+1)
//	)
//	for i := range dp {
//		dp[i] = make([]int, len2+1)
//	}
//	for i := 1; i <= len1; i++ {
//		for j := 1; j <= len2; j++ {
//			if s1[i-1] == s2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//			}
//		}
//	}
//	return asciiSum(s1) + asciiSum(s2) - dp[len1][len2]*2
//}
//
//
//



