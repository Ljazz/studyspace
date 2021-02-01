package main

import (
	"fmt"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
	name string
}

// newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student{
	return &student{
		id: id,
		name: name
	}
}

func showAllStudent() {
	// 遍历 把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s", k, v.name)
	}
}

func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1. 创建一个新学生
	// 1.1 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	// 1.2 调用构造函数
	newStu := newStudent(id, name)
	// 2. 追加到allStudent这个map中
	allStudent[id] = newStu
}

func deleteStudent() {
	// 请用户输入要删除学生的学号
	var(
		deleteId int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&deleteId)
	// 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int64]*student, 50)
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统!")
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出
	`)
		fmt.Print("请输入您的操作：")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("您选择了%d这个选项！\n", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("滚~")
		}
	}
}
