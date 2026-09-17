package main

import "fmt"

func main() {
	var pointer *int

	fmt.Println(pointer)

	if pointer != nil {
		fmt.Println("Not nill")
	}

	value := 42
	pointer = &value
	fmt.Println(pointer)
	fmt.Println(*pointer)

	*pointer += 10
	fmt.Println(*pointer)
}
