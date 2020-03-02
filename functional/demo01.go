package main

import "fmt"

func main() {
	// https://www.flysnow.org/2017/03/31/go-in-action-go-method.html
	// 在Go语言中，函数是指不属于任何结构体、类型的方法，也就是说，函数是没有接收者的；
	sum := add(1, 2)
	fmt.Println(sum)

	// 方法有两种类型的接收者：值接收者和指针接收者
	v := person{name: "张三"}
	p := person{name: "改变姓名"}
	v.modify() // 值接收者，修改无效
	fmt.Println(p.String())

}

// 小写开头私有
func add(a, b int) int {
	return a + b
}

type person struct {
	name string
}

func (p person) String() string {
	return "the person name is " + p.name
}

func (p person) modify() {
	p.name = "李四"
}

func (p *person) modifyP() {
	p.name = "李四"
}
