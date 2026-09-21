package main

import (
	"fmt"
)

func main() {
	var zeroSlice []int
	fmt.Println(zeroSlice, zeroSlice == nil, len(zeroSlice))

	var oneElementSlice = make([]int, 1, 10)
	fmt.Println(oneElementSlice, len(oneElementSlice), cap(oneElementSlice))

	sliceWithLenEqCap := make([]int, 5)
	fmt.Println(sliceWithLenEqCap, len(sliceWithLenEqCap), cap(sliceWithLenEqCap))

	myFavSlice := []string{"I", "like", "learning", "Go"}
	fmt.Println(myFavSlice)

	fmt.Println(myFavSlice[1:2])
	fmt.Println(myFavSlice[:3])

	myArr := [5]int{20, 15, 5, 30, 25}
	mySlice := myArr[1:4]

	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", myArr, len(myArr), cap(myArr))
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	for _, v := range myFavSlice {
		fmt.Println(v)
	}

	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))
	e := copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(e)

	s3 := make([]string, 2)
	e2 := copy(s3, s1)
	fmt.Println(s3, e2)

	foo := []int{1, 2}
	bar := []int{}
	copy(bar, foo)
	fmt.Println(bar)
}
