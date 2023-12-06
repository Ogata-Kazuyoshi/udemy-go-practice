package main

import "fmt"


func main()  {
	f := func (x int, y int) int  {
		return x + y
	}
	i := f(1,2)
	fmt.Println(i)
	fmt.Printf("%T\n", f)
	
}