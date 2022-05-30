package main

import (
	"fmt"

	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	date.Month = 12
	date.SetYear(2021)
	fmt.Println(date)
	fmt.Printf("%v\n", date)
	fmt.Printf("%#v\n", date)

}
