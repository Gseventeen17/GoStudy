package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内变量
//var aaa = 100
//var bbb = "hhh"
//不可用:=

const filename = "a.txt"

func consts() {
	const (
		a = 5
	)
}

// 枚举，一般也定义为const
/*func enums() {
	const (
		java   = 0
		python = 1
		cpp    = 2
	)
	fmt.Println(java, python, cpp)
}*/

func enums() {
	const (
		//iota 自增值
		java = iota
		python
		cpp
	)

	//iota可以参与运算
	const (
		zero    = 10 << iota //10*(2的0次方)
		ten                  //10*(2的1次方)
		hundred              //10*(2的2次方)
	)

	fmt.Println(zero, ten, hundred)

	fmt.Println(java, python, cpp)
}

var (
	aaa = 100
	bbb = "hhh"
)

func variableZero() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitial() {
	//var c, d int = 1, 2
	var a, b = 4, 5
	//var str string= "abc"
	var str = "abc"
	println(a, b, str, aaa, bbb)
}

func variableTypeDeduction() {
	var a, b, c = 1, true, "你好"
	println(a, b, c)
}

func variableShorter() {
	a, b, c := 1, true, "你好"
	println(a, b, c)
}

func eular() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triAngle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("hello world")
	variableZero()
	variableInitial()
	variableTypeDeduction()
	variableShorter()
	eular()
	triAngle()
	enums()
}
