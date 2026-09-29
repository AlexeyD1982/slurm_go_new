package main

import (
	"fmt"
)

func echo(arg string, ch chan string) {
	ch <- arg
	fmt.Println("echo bye")
}

func main() {
	var ch chan string
	fmt.Println(ch)

	chInit := make(chan string)
	fmt.Println(chInit)

	go echo("Hello world", chInit)
	fmt.Println("Bye0")
	data := <-chInit
	fmt.Println(data)
	fmt.Println("Bye")

	//buffered channels
	ch1 := make(chan string, 2)
	go echo("value 1", ch1)
	go echo("value 2", ch1)
	val1 := <-ch1
	val2 := <-ch1
	fmt.Println("finish")
	fmt.Println(val1, val2)

	close(ch1)

	var c = make(chan int, 3)
	c <- 20
	c <- 10
	c <- 0
	close(c)
	for i := 0; i < 5; i++ {
		v, ok := <-c
		fmt.Printf("closed?: %v, value %d\n", !ok, v)
	}

	var c2 = make(chan int, 3)
	c2 <- 20
	c2 <- 10
	c2 <- 0
	close(c2)
	for v := range c2 {
		fmt.Printf("Value: %d\n", v)
	}
}

func directedEcho(arg string, ch chan<- string) {
	ch <- arg
}

func directedReceive(ch <-chan string) {
	<-ch
}

func directedReceiveIfClosed(ch <-chan int) {
	value, ok := <-ch
	if ok {
		fmt.Println("process value", value)
	}
}
