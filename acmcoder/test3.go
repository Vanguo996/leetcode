package main

import "fmt"

type node struct {
	b byte
	num int
}

func sort(arr []node, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	sort(arr, l, p - 1)
	sort(arr, p + 1, r)
}

func partition(arr []node, l, r int) int{
	pivot := arr[r].num
	for i := l; i < r; i++{
		if arr[i].num < pivot{
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}
	arr[l], arr[r] = arr[r], arr[l]
	return l
}

func sort2(arr []node) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j].b < arr[j + 1].b {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func count(str string) {
	c := [256]int{}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || (str[i] >= '0' && str[i] <= '9') ||
			(str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') {
			c[str[i]]++
		}
	}

	//装入到node切片中
	nodes := []node{}
	for i, v := range c {
		if v != 0 {
			nodes = append(nodes, node{b: byte(i) , num: v})
		}
	}
	//按照频率排序，
	sort(nodes, 0, len(nodes) - 1)

	//对相同频率的ascii码排序
	r,l := 0,0
	for i := 0; i < len(nodes) - 1; i++ {
		if nodes[i].num != nodes[i + 1].num {
			r = i
			if l != r {
				sort2(nodes[l + 1 : r + 1])
			}
			l = i
		} else {
			r++
		}
	}

	for _, v := range nodes {
		fmt.Println(v.b, v.num)
	}
}

func main() {
	//count("accaacadvdgds2346fdv23421241sdvsdaf")
}



func reverse() {

}



