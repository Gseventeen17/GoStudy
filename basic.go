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
}
