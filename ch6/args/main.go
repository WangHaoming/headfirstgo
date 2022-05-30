package main

import (
	"fmt"
)

func main() {
	// sliceTest := []int{1, 2, 3}
	// meke生成的切片不是一个空指针（nil），但是数组本身是一个控数组
	// sliceTest2 := make([]int, 0)

	// 变量声明只生成一个空指针
	var sliceTest2 []int
	fmt.Printf("%T\n", sliceTest2)
	fmt.Println(sliceTest2 == nil)

	// append函数既可以接受空数组，也可以接受一个空指针（nil）
	sliceTest2 = append(sliceTest2, 3)
	// sliceTest2 = append(sliceTest2, 3, 4, 5)
	// fmt.Println(sliceTest[4])
	fmt.Println(sliceTest2)
}
