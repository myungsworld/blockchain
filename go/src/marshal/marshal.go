package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User2 struct {
	// 미출력
	Name string `json:"-"`
	// 문자열로 출력
	Ages int `json:"age, string"`
	// 값이 비어있으면 미출력
	Phone string `json:omitempty`
}

func main() {
	//var u = User{"Go", 7}
	//b, _ := json.Marshal(u)

	//fmt.Println(b)
	//fmt.Println(string(b))

	var u2 = User2{"song", 10, ""}
	b2, _ := json.Marshal(u2)
	fmt.Println(b2)
	fmt.Println(string(b2))

	var s = `{"name":"song","age":27}`
	var u User
	json.Unmarshal([]byte(s), &u)
	fmt.Printf("%+v\n", u)
}
