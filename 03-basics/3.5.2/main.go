package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(frequentWord("ad t pi mk ht ad pi t ad"))
}

func frequentWord(str string) string {
	strArr := strings.Split(str, " ")
	var maxWord string
	var mx int
	mp := make(map[string]int)

	for _, word := range strArr {
		var newCol int
		if col, ok := mp[word]; ok {
			newCol = col + 1
			mp[word] = newCol
		} else {
			newCol = 1
			mp[word] = newCol
		}
		if newCol > mx {
			mx = newCol
			maxWord = word
		}
	}

	return maxWord
}
