package main

import (
	"fmt"
	"time"
)

func echo(arg string) {
	fmt.Println(arg)
}

func main() {
	go echo("hello")
	time.Sleep(1 * time.Millisecond)
}
