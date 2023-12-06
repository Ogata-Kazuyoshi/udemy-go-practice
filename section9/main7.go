package main

import (
	"fmt"
)

func double(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	var p *int = &n
	double(p)
	fmt.Println(n)
}
