// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		// log.Fatal(err)
	}
	grade1 := int32(grade)
	fmt.Println(reflect.TypeOf(grade1))
	fmt.Println(grade1)
}
