package main

import (
	"encoding/json"
	"fmt"
	"os"
)
func main()  {
	file, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //deferなので後で閉じられる。開いたら閉じるを先に書いとくのも良いね
	file.Write([]byte("Hello Goです"))
	file2, err2 := os.Create("test.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file2.Close()

	data := map[string]int{
		"apple": 100,
		"banana": 200,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	file2.Write(jsonData)
}