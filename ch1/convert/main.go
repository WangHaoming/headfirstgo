package main

import "fmt"

func main() {
	var length float64 = 1.2
	var x interface{}
	var myInt int = 2
	x = myInt
	fmt.Println(length)
	length, ok := x.(float64)
	if !ok {
		fmt.Println("convert failed!")
	}
	fmt.Println(length)
	length = float64(myInt)
	fmt.Println(length)
	switch x.(type) {
	case int:
		fmt.Println("the type of a is int")
	case float64:
		fmt.Println("the type of a is float")

	}
}
