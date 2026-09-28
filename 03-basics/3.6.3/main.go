package homework

import "unicode"

func removeSpaces(s string) string {
	resArr := []rune{}
	for _, r := range s {
		if !unicode.IsSpace(r) {
			resArr = append(resArr, r)
		}
	}
	return string(resArr)
}
