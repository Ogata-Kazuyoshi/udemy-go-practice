package main

import (
	"fmt"
	"time"
)

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)
	// ch1 <-1
	// close(ch1)

	// i,ok := <-ch1
	// fmt.Println(i,ok)
	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 10 {
		ch1 <- i
		time.Sleep(100 * time.Millisecond)
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
