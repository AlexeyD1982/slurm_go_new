package main

import "fmt"

func main() {
	fmt.Println(factorialRec(8))
	fmt.Println(factorialIter(8))
	//findLimit(1)
}

func factorialRec(n int) int {
	if n < 2 {
		return n
	}
	return n * factorialRec(n-1)
}

func factorialIter(n int) int {
	res := 1
	for x := n; x > 1; x-- {
		res *= x
	}
	return res
}

func findLimit(n int) {
	fmt.Println(n)
	findLimit(n + 1)
}
