package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	m := make(map[string]string)

	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key:", key, "value", value)
	}

	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	p := person{name: "song", age: 27}
	fmt.Println(p)

	i := 7
	inc(&i)
	fmt.Println(i)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}
