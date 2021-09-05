package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {

	wordFreq("how do you do")

}

func wordFreq(s string) {
	s1 := strings.Split(s, " ")
	m := make(map[string]int)
	for _, str := range s1{
		if m[str] == 0 {
			m[str] = 1
		} else {
			m[str] = m[str] + 1
		}
	}

	for k, v := range m{
		fmt.Println(k ,v)
	}
}

func mapDemo() {


	//m := make(map[string]int, 10)
	//
	//m["van"] = 24
	//
	//for v, ok := range m{
	//	fmt.Print(v, ok)
	//}
	//delete(m, "van")

	rand.Seed(time.Now().Unix())

	m := make(map[string]int)

	for i := 0; i < 100; i++{
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		m[key] = value
	}

	//创建一个string切片
	var keys = make([]string, 0, 100)
	for key := range m{
		keys = append(keys, key)
	}

	//排序输出key
	sort.Strings(keys)
	for _, key := range keys{
		fmt.Println(key, m[key])
	}

}

func SliceOfMapType() {
	//创建一个map类型的切片
	//var slice = make([]map[string]int, 3)
	//for index, v := range slice{
	//
	//}
}

func palidrome(s string) {

	//for i, v := range s{
	//	if s[i] != s[len(s) - 1]{
	//
	//	} else{
	//
	//	}
	//}

	for i := 0; i < len(s) / 2; i++{

	}

}

func deferDemo() int{
	x := 5
	defer func(){
		x++
	}()
	return x
}




