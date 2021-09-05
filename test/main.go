package main

import "fmt"

//func main() {
//
//	//通过构造函数
//	p1 := newPerson("小王子", 25)
//	p1.Dream()
//
//
//}

//func NewPerson(s string, i int) interface{} {
//
//}

//type person struct {
//	name string
//	age  int8
//}
//
//func newPerson(name string, age int8) *person {
//	return &person{
//		name: name,
//		age:  age,
//	}
//}
////隐式的实现接口
////显式的实现接口  java  implement interface
//
////值类型的接收者
//func (this person) Dream() {
//
//}
//

// Sayer 接口
type Sayer interface {
	say()

}

type Mover interface {
	move()

}


type dog struct {}

type cat struct {}

func (d dog) say() {
	fmt.Print("sss")
}




func main() {

}

