package main

import "fmt"

func main() {
	var arr [4]int
	fmt.Println(arr[3])

	var arr2 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr2)

	arr3 := [4]int{10, 20, 30, 40}
	fmt.Println(arr3)

	multiDimArr := [...][4]int{
		{1, 2, 3, 5},
		{10, 20, 30, 40},
	}
	fmt.Println(multiDimArr)

	a1 := [3]string{"a", "b", "c"}
	var b1 = a1
	b1[1] = "z"
	fmt.Println(a1)
	fmt.Println(b1)
}
