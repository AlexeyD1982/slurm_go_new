package main

import "fmt"

func main() {
	ranges()
	typeEmbedding()
	valueReceivers()
	nilProblem()
	stringImmutable()
	staleSlices()
}

func ranges() {
	valuesStr := []string{"a", "b", "c"}
	for index, value := range valuesStr {
		fmt.Println(index, value)
	}

	valuesInt := []int{4, 8, 23}
	for value := range valuesInt {
		fmt.Println(value)
	}
}

func typeEmbedding() {
	v1 := Wrapper{Node{"123", 0}}
	fmt.Println(v1.Name(), v1.Node.Name())
	fmt.Println(v1.value, v1.Node.value)

}

type Node struct {
	value  string
	weight int
}

func (n Node) Name() string {
	return n.value
}

type Wrapper struct {
	Node
}

func (w Wrapper) Name() string {
	return "wrapped value" + w.Node.value
}

func valueReceivers() {
	n := Node{value: "my value"}
	n.SetValue("new value")
	fmt.Println(n.value)
}

func (n Node) SetValue(newValue string) {
	n.value = newValue
}

func nilProblem() {
	var slice []int
	fmt.Println(len(slice))
	slice = append(slice, 100500)
	fmt.Println(len(slice))

	var m map[string]int
	fmt.Println(len(m))
	//m["test"] = 1

	mOk := make(map[string]int)
	fmt.Println(len(mOk))
	mOk["one"] = 1
	fmt.Println(len(mOk))
}

func stringImmutable() {
	s := "abc"
	//s[1] = 'B'

	sBytes := []byte(s)
	sBytes[1] = 'B'
	fmt.Println(string(sBytes))

	sRunes := []rune(s)
	sRunes[1] = 'B'
	fmt.Println(string(sRunes))

	s2 := "ЩЫ"
	r2 := []rune(s2)
	fmt.Println(string(r2))
	r2[0] = 'S'
	fmt.Println(string(r2))
}

func staleSlices() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2)

	s2[0] = 100500

	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 4)
	//s2[0] = 10
	fmt.Println(s1)
	fmt.Println(s2)

}
