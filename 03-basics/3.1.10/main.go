package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m, m == nil)
	//m["test"] = 1

	var m2 = make(map[string]int)
	fmt.Println(m2, m2 == nil)
	m2["test"] = 1
	m2["test1"] = 2
	fmt.Println(m2)

	var m3 = map[string]int{
		"hi": 100,
	}
	fmt.Println(m3)

	v, ok := m2["test2"]
	if !ok {
		fmt.Println("Value doesn't exist")
	} else {
		fmt.Println(v)
	}

	delete(m2, "test2")

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m5 := m2
	m5["new"] = 4
	fmt.Println(m5, m2)
}
