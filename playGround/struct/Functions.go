package main

import (
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
}



func main() {
	for i := 0; i < 100; i++{
		go func(i int) {
			fmt.Printf("hello %d\n", i)  //
		}(i)
	}
	fmt.Println("end")
}



type student struct {
	name string
	age  int
}

func test() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus{
		m[stu.name] = &stu  //取值
	}

	for k, v := range m{
		fmt.Println(k, "=>", v.name)
	}

	//小王子 => 大王八
	//娜扎 => 大王八
	//大王八 => 大王八

}


