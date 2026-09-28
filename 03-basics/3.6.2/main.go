package main

import "fmt"

func main() {
	fmt.Println(reverse("hello"))
}

func reverse(s string) string {
	arr := []rune(s)
	resArr := make([]rune, len(arr))
	for i, r := range arr {
		resArr[len(arr)-1-i] = r
	}
	return string(resArr)
}
