package main

import (
	"fmt"
)

func main() {
	s4 := make(map[string]int)
	s4["Apple"] = 100
	s4["Banana"] = 200
	s, flg := s4["Apple"]
	if !flg {
		fmt.Println(flg)
	} else {
		fmt.Println(s)
	}
}
