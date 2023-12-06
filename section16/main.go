package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"` //`json:`はJSONにした際のキーになる
	Name    string    `json:"name"` //uのオブジェクトの中にはあるが、隠したい場合は、"name,omitempty"としてあげる
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct{
		Name string
	}{
		Name : "Mr. " + u.Name,
	})
	return v,err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@gmail.com"
	u.Created = time.Now()

	//構造体(オブジェクト)からjsonに変換
	// marshal jsonに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	//jsonから構造体に変換
	u2 := new(User)

	//Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs,&u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
