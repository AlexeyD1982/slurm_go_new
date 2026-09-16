package main

import "fmt"

var packageLevelVar = "I'm outside function"

func main() {
	fmt.Println(packageLevelVar)

	var str string
	var strWithInitialnization string = "I'm initialized string"
	var str1, str2 string
	fmt.Println(strWithInitialnization, str, str1, str2)

	var strWithoutExplicitType = "I'm string data type"
	strShorthandDeclaration := "I'm string with shorthand declaration"
	fmt.Println(strWithoutExplicitType, strShorthandDeclaration)

	const myStr = "I'm string constant"
	fmt.Println("myStr[0]=", myStr[0])
	var multiLineStr string = `I'm
multiline
string`
	fmt.Println(multiLineStr)
}
