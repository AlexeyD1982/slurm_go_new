package main

import (
	"fmt"
	"time"
)

func main() {
	//selectExample()
	//selectExampleDefault()
	//severalChannels()
	//timeout()
	finishWork()
}

func selectExample() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		one <- "one"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		one <- "two"
	}()

	select {
	case result := <-one:
		fmt.Println("Received:", result)
	case result := <-two:
		fmt.Println("Received:", result)
	}

	close(one)
	close(two)
}

func selectExampleDefault() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		one <- "One"
	}()

	go func() {
		two <- "Two"
	}()

	select {
	case result := <-one:
		fmt.Println("Received:", result)
	case result := <-two:
		fmt.Println("Received:", result)
	default:
		fmt.Println("Default...")
	}

	close(one)
	close(two)
	//time.Sleep(100 * time.Millisecond)
}

func severalChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}

func timeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "data received"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: no data received within 2 seconds")
	}
}

func finishWork() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("Sent data to channel")
	default:
		fmt.Println("Channel is full, cannot send data")
	}

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No data available")
	}

}
