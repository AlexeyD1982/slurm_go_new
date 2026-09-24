package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(add(2147483647))
}

func add(a int32) int32 {
	if a < math.MaxInt32 {
		return a + 1
	}
	return a
}
