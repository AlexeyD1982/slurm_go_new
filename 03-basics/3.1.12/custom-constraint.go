package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyConstraint interface {
	Value() string
}

func test[T MyConstraint](x T) {
	fmt.Println(x.Value())
}

type MyStruct1 struct {
	value string
	a     int
}

func (x MyStruct1) Value() string {
	return x.value
}

type MyStruct2 struct {
	value string
	b     float32
}

func (x MyStruct2) Value() string {
	return x.value
}

type myString string

type MyConstraintAllStrings interface {
	constraints.Integer | ~string
}

type MyConstraintStrictStrings interface {
	constraints.Integer | string
}

func allStrings[T MyConstraintAllStrings](x T) {
	fmt.Println(x)
}

func strictStrings[T MyConstraintStrictStrings](x T) {
	fmt.Println(x)
}

type Number interface {
	int | int8 | int16
}

type Vector[T Number] []T

func addVectors[T Number](vec1 Vector[T], vec2 Vector[T]) Vector[T] {
	var result Vector[T]
	for i := range vec1 {
		result = append(result, vec1[i]+vec2[i])
	}
	return result
}

func customConstraints() {
	test(MyStruct1{value: "test1"})
	test(MyStruct2{value: "test2"})

	allStrings("testAll")
	allStrings(myString("testAll"))
	strictStrings("testStrict")
	//strictStrings(myString("testStrict"))

	v1 := Vector[int]{1, 2, 3}
	v2 := Vector[int]{3, 4, 5}
	fmt.Printf("addVectors(v1, v2)=%f\n", addVectors(v1, v2))
}
