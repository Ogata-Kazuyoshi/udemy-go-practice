package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	ch1 <- 10
	ch2 <- "A"
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 100)
	case v2 := <-ch2:
		fmt.Println(v2 + "!")
	default:
		fmt.Println("どちらでもないです。")
	}
}
