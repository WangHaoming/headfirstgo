package main

// 从headfirstgo/ch2/util导入utila包
// 同时也可以表示从headfirstgo/ch2/util导入util包，并重命名包名为utila
import (
	"fmt"
	utila "headfirstgo/ch2/util"
)

func main() {
	err := utila.PackageTest()
	fmt.Println(err.(error))
	fmt.Println(err)
}
