package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "Streёt"
	fmt.Println(len(word))
	fmt.Println(word[4:])
	fmt.Println(utf8.RuneCountInString(word))

	for pos, char := range word {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
}
