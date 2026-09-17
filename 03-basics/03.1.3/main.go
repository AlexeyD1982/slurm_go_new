package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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

	const a1 = 10
	const a2 = a1

	var b1 = 10
	fmt.Println(b1)

	//const b2 = b1

	var myBool1 = rand.Int() > 10
	var myBool2 = rand.Int() < 10
	var iAmTrueOrFalse = myBool1 || myBool2
	_, _ = multiLineStr, iAmTrueOrFalse

	var i int = 404
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	_, _, _, _, _ = i, i8, i16, i32, i64

	var ui uint = 404
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	_, _, _, _, _ = ui, ui8, ui16, ui32, ui64

	var f32 float32 = 1.7812
	var f64 float64 = 3.1415
	_, _ = f32, f64

	mainQuestion := 42
	f := float64(mainQuestion)
	u := uint(f)
	_ = u

	iAmInt64, _ := strconv.ParseInt("100500", 10, 64)
	fmt.Printf("%d, %T\n", iAmInt64, iAmInt64)
	iAmInt64Overflow, err := strconv.ParseInt("100500", 10, 8)
	fmt.Println(iAmInt64Overflow, err)

	if x := rand.Int(); x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is greater than or equal to 5")
	}

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
		fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday")
	case "Wednesday":
		fmt.Println("It's Wednesday")
	default:
		fmt.Println("It's not Monday, Tuesday, or Wednesday")
	}

	x := rand.Int()

	switch {
	case x > 5:
		fmt.Println("x is greater than 5")
		fmt.Printf("%d, %T\n", x, x)
	default:
		fmt.Println("x is less than or equal to 5")
		fmt.Printf("%d, %T\n", x, x)
	}

	fClosure := myFuncWithClosure()
	fClosure(10)
	fmt.Println(fClosure(20))

	fmt.Println(add(1, 2, 3, 4, 5))

	closureErr()

	for i := range 5 {
		fmt.Println(i)
	}
}

func myFunctionFirstClass() {
	fn := func() {
		fmt.Println("Hello, world!")
	}
	fn()
}

func myFunctionAnon() {
	func() {
		fmt.Println("inside fu")
	}()
}

func myFuncWithClosure() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v

		return sum
	}
}

func add(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}
	return sum
}

func closureErr() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}
}
