package main

import (
	"fmt"
)

func foo(v []int) []int {
	v = append(v, 100)
	return v
}

func main() {
	v := make([]int, 5, 10)
	foo(v)
	fmt.Println(v)
}
