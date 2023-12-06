package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	go reciever(ch1)
	go reciever(ch2)
	i := 0
	for i < 10 {
		ch1 <- i
		ch2 <- i * 100
		time.Sleep(2000 * time.Millisecond)
		i++
	}
}
