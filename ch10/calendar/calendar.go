package calendar

type Date struct {
	year  int
	Month int
	Day   int
}

// set方法需要一个指针类型的接收器，因为需要改变结构体变量内部的属性
func (d *Date) SetYear(year int) { d.year = year }
