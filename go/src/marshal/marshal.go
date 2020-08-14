package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b, _ := json.Marshal(true)
	fmt.Println(b)
	fmt.Println(string(b))
}
