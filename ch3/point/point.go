package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFlout float64
	fmt.Println(reflect.TypeOf(&myFlout))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
}
