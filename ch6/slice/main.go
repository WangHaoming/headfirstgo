package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(underlyingArray)
	slice1 := underlyingArray[0:6]
	fmt.Println(slice1)
}
