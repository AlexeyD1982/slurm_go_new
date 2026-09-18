package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

func (c Person) doNothing() {}

func (c Person) CanDriveCar() bool {
	return c.Birthday.AddDate(18, 0, 0).Before(time.Now())
}

func (c Person) UpdateName(name string) {
	c.Name = name
}

func (c *Person) UpdateNameByPointer(name string) {
	c.Name = name
}

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	const layout = "2006-Jan-02"
	tm, _ := time.Parse(layout, "2003-Sep-03")
	p := Person{
		Name:     "John",
		Birthday: tm,
	}
	fmt.Println(p)
	fmt.Println(p.CanDriveCar())
	p.UpdateName("Jim")
	fmt.Println(p)
	p.UpdateNameByPointer("Jared")
	fmt.Println(p)

	i := MyInt(10)
	fmt.Println(i.isGreater(10))
}
