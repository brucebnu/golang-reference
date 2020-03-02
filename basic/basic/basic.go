package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// package 作用域包内变量，没有全局变量说法
// 括号省略多个var关键词
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def" //编译器自动识别变量两类型
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 变量定义一定用到，不用不行；
	// 第一次定义变量用冒号，第二次赋值不用，只能在函数内使用；
	//
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/**
欧拉公式
*/
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

// 三角函数
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 常量定义
// 可以定义函数、包内
// 不建议大写，大小写有特殊含义，首字母大写表示public
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 特殊常量，枚举类型
func enums() {
	const (
		cpp = iota // iota 表示下面变量自增
		_          // 自增跳过
		python
		golang
		javascript
	)

	// 高级iota用法
	//
	const (
		b = 1 << (10 * iota) // 1左移，10乘以iota位
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
