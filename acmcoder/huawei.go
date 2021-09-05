package main

import (
	"fmt"
)

//func main(){
//	input := bufio.NewScanner(os.Stdin)
//	//input.Scan()
//	//s := input.Text()
//	//solution(s)
//
//	input.Scan()
//	//a := input.Text()
//	//b, _ := strconv.Atoi(a)
//
//	input.Scan()
//	s := input.Text()
//
//	s1 := strings.Split(s, " ")
//
//	fmt.Println(s1)
//	//fmt.Print(reflect.TypeOf(s))
//
//	arr := []int{}
//	for _, v := range s{
//		if v != ' ' {
//			fmt.Print(v)
//			//arr = append(arr, a)
//		}
//	}
//
//	solution2(arr)
//
//
//}

type student struct {
	name string
	age int
}

func main() {
	m := make(map[string]*student)

	stu := []student{
		{name: "a", age: 1},
		{name: "b", age: 2},
	}

	for _, v := range stu{

		m[v.name] = &v

	}

	for k, v := range m {
		fmt.Println(k, v)
	}


}

func solution(s string) {
	fmt.Println(s)
}

func solution2(nums []int) {
	fmt.Print(nums)
}

