package main

import "fmt"

type NewInt int
type NewInt2 int

func main() {
	myInt := NewInt(10)
	// 两个自定一类型不能相加（不能做运算，即使他们的基础类型相同）
	// myInt = myInt + NewInt2(10)
	// 但是自定一类型可以和基础类型相加,一个定义的类型支持与基础类型相同的运算符
	myInt = myInt + 10
	fmt.Printf("%T\n", myInt)
	fmt.Println(myInt)
}
