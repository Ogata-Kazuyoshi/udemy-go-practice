package main

import (
	"fmt"
)

func main() {

	// var ch1 chan int
	// ch1 = make(chan int)

	// ch2 := make(chan int)

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println(ch3)
	fmt.Println(len(ch3))
	i := <-ch3
	fmt.Println("i : ", i)
	fmt.Println("len : ", len(ch3))
	i2 := <-ch3
	fmt.Println("i2 : ", i2)
	fmt.Println("len : ", len(ch3))

	fmt.Println("i3 : ", <-ch3)
	fmt.Println("len : ", len(ch3))

	// //受診せ引用
	// var ch2 <-chan int

	// //送信専用
	// var ch3 chan<- int

}
