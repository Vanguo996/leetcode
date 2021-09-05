package main

import "fmt"

type student struct {
	id int64
	name string
}

// 学生管理
type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) add() {

	var(
		id int64
		name string
	)

	//获取用户输入
	fmt.Println("输入id")
	fmt.Scanln(&id)
	fmt.Println("输入name")
	fmt.Scanln(&name)
	newStu := student{
		id:   id,
		name: name,
	}
	// allStudent本质上是一个map
	s.allStudent[newStu.id] = newStu

}

func (s studentMgr) edit() {
	fmt.Println("输入修改的学号：")


}

func (s studentMgr) show() {
	for _, stu := range s.allStudent{
		fmt.Printf("学号: %d  姓名：%s", stu.id, stu.name)
		fmt.Println( " ")
	}

}

func showMenu() {
	fmt.Println("welcome sms")
}

func stuManager(){
	smr := studentMgr{allStudent: make(map[int64]student, 100),}

	for {
		showMenu()
		fmt.Print("输入序号")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("输入为：", choice)
		switch choice {
		case 1:
			smr.add()
		case 2:
			smr.show()
		case 3:
			smr.edit()
		default:
			fmt.Print("default")

		}
	}
}

